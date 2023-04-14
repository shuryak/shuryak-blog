import React, {useState, useEffect, useRef} from 'react'
import {OutputData} from '@editorjs/editorjs'
import '@styles/editor.scss'
import {EditorBlock} from '../components/EditorBlock'
import PublicationBlock from '../components/PublicationBlock'
import {ShortArticle} from '../models/dto'
import {getArticles} from '../api/api'

export const EditorPage: React.FC = () => {
  const [data, setData] = useState<OutputData>()
  const [isPublish, setIsPublish] = useState<boolean>(false)
  const [drafts, setDrafts] = useState<ShortArticle[]>([])

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
    console.log(data)
  }, [data])

  return (
    <>
      <div className="page-header-with-drafts">
        <h1>{'📝 Редактор' + (isPublish ? ': сохранение' : '')}</h1>
        { !isPublish &&
        <form className="drafts-form">
          <div className="input-block">
            <select className="account-form-input user-input" name="" id="drafts">
              {drafts.map((draft, idx) =>
                <option value={draft.id} key={idx}>
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
          <EditorBlock data={data} onChange={setData} holder="editorjs"/>
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
