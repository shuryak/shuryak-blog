import React, {useState, useEffect, useRef} from 'react'
import {OutputData} from '@editorjs/editorjs'
import '@styles/editor.scss'
import PublicationBlock from '../components/PublicationBlock'
import {ShortArticle} from '../models/dto'
import {getArticles, createArticle, updateArticle, getArticleByCustomId} from '../api/api'
import {v4 as uuidv4} from 'uuid'
import {useDispatch, useSelector} from 'react-redux'
import {RootState} from '../store/reducers/rootReducer'
import {createReactEditorJS} from 'react-editor-js'
import {EDITOR_JS_TOOLS} from './editor-js-tools'
import {EditorCore} from '@react-editor-js/core/src/editor-core'

export const EditorPage: React.FC = () => {
  const [data, setData] = useState<OutputData>()
  const [isPublish, setIsPublish] = useState<boolean>(false)
  const [drafts, setDrafts] = useState<ShortArticle[]>([])
  const [draftArticle, setDraftArticle] = useState<ShortArticle | null>(null)

  const {accessToken} = useSelector(
    (state: RootState) => state.login
  )

  const calledOnce = useRef(false);

  useEffect(() => {
    if (calledOnce.current) {
      return;
    }

    getArticles({
      count: 10,
      offset: 0,
      isDrafts: true
    }).then(data => {
      setDrafts(data.data.articles)
    })

    calledOnce.current = true;
  }, [])

  useEffect(() => {
    if (!draftArticle && accessToken) {
      createArticle({
        customId: uuidv4(),
        title: `Черновик от ${new Date().toLocaleDateString('ru-ru',
          {
            year: 'numeric',
            month: 'short',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric'
          })}`,
        thumbnail: 'https://sun9-10.userapi.com/impg/EyfFwRqAPqTv2Ody1bkD17RfyUDpWXrZ47wgnQ/TEoD8QO9Z2o.jpg?size=1280x961&quality=95&sign=099dd9caf8c90583d8d66cbf3cd197fd&type=album',
        content: data!,
        isDraft: true
      }, accessToken!).then(data => {
        setDraftArticle(data.data)
      })
    } else if (accessToken) {
      updateArticle({
        customId: draftArticle!.customId,
        title: draftArticle!.title,
        thumbnail: draftArticle!.thumbnail,
        content: data!,
        isDraft: true
      }, accessToken!).then(data => {
        setDraftArticle(data.data)
      })

      getArticles({
        count: 10,
        offset: 0,
        isDrafts: true
      }).then(data => {
        setDrafts(data.data.articles)
      })
    }
  }, [data])

  const loadData = (e: React.ChangeEvent<HTMLSelectElement>) => {
    getArticleByCustomId(e.target.value).then(data => {
      setData(data.data.content)
      editorCore.current?.render(data.data.content)
      setDraftArticle({
        authorId: data.data.authorId,
        createdAt: data.data.createdAt,
        customId: data.data.customId,
        id: data.data.id,
        isDraft: data.data.isDraft,
        shortContent: '',
        thumbnail: data.data.thumbnail,
        title: data.data.title,
        updatedAt: data.data.updatedAt
      })
    })
  }

  const ReactEditorJS = createReactEditorJS()

  const editorCore = React.useRef<EditorCore | null>(null)
  const handleInitialize = React.useCallback((instance: any) => {
    editorCore.current = instance
  }, [])

  const handleSave = React.useCallback(async () => {
    if (editorCore.current) {
      const savedData = await editorCore.current.save()
      setData(savedData)
    }
  }, [])

  return (
    <>
      <div className="page-header-with-drafts">
        <h1>{'📝 Редактор' + (isPublish ? ': сохранение' : '')}</h1>
        { !isPublish &&
        <form className="drafts-form">
          <div className="input-block">
            <select className="account-form-input user-input" name="" id="drafts" onChange={loadData}>
              {drafts?.map((draft, idx) =>
                <option value={draft.customId} key={idx}>
                  {draft.title}
                </option>
              )}
            </select>
          </div>
        </form>
        }
      </div>
      {!isPublish ?
        <>
          <ReactEditorJS onChange={handleSave} onInitialize={handleInitialize} holder="editorjs" tools={EDITOR_JS_TOOLS}>
            <div id="editorjs"></div>
          </ReactEditorJS>
          <button className="default-btn user-input" onClick={() => setIsPublish(true)}>
            Сохранить
          </button>
        </>
        :
        <PublicationBlock data={data} goBack={() => setIsPublish(false)}/>
      }
    </>
  )
}

export default EditorPage
