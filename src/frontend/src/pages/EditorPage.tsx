import React, {useState, useEffect} from 'react'
import {OutputData} from '@editorjs/editorjs'
import '@styles/editor.scss'
import {EditorBlock} from '../components/EditorBlock'
import PublicationBlock from '../components/PublicationBlock'

export const EditorPage: React.FC = () => {
  const [data, setData] = useState<OutputData>()
  const [isPublish, setIsPublish] = useState<boolean>(false)

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
              <option value="">Новый черновик</option>
              <option value="">Lorem ipsum dolor.</option>
              <option value="">Lorem ipsum dolor 2.</option>
              <option value="">Lorem ipsum dolor 3.</option>
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
