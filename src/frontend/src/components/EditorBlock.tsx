import React, {memo, useEffect, useRef} from 'react'
import EditorJS, {ToolConstructable, OutputData} from '@editorjs/editorjs'
import Header from '@editorjs/header'
// @ts-ignore
import List from '@editorjs/list';
// @ts-ignore
import NestedList from '@editorjs/nested-list';
// @ts-ignore
import CheckList from '@editorjs/checklist';
// @ts-ignore
import Code from '@editorjs/code';
// @ts-ignore
import Marker from '@editorjs/marker';
// @ts-ignore
import InlineCode from '@editorjs/inline-code';
// @ts-ignore
import HyperLink from 'editorjs-hyperlink';
import '@styles/editor.scss'

export const EditorBlock: React.FC<{data: OutputData | undefined, onChange: (data: OutputData) => void, holder: string}> = ({data, onChange, holder}) => {
  const ref = useRef<EditorJS>()

  useEffect(() => {
    if (!ref.current) {
      const editor = new EditorJS({
        holder,
        data,
        tools: {
          header: {
            class: Header as unknown as ToolConstructable, // TODO: Header type fix
            config: {
              placeholder: 'Текст заголовка',
              levels: [1, 2, 3]
            }
          },
          List,
          NestedList,
          CheckList,
          Code,
          Marker,
          InlineCode,
          HyperLink
        },
        async onChange(api, event) {
          const data = await api.saver.save()
          onChange(data)
        }
      })
      ref.current = editor
    }

    return () => {
      if (ref.current && ref.current.destroy) {
        ref.current.destroy()
      }
    }
  },[])

  return <div id={holder}></div>
}

export default memo(EditorBlock)
