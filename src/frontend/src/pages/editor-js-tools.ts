import Header from '@editorjs/header'
// @ts-ignore
import List from '@editorjs/list';
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
import {ToolConstructable} from '@editorjs/editorjs'

export const EDITOR_JS_TOOLS = {
  header: {
    class: Header as unknown as ToolConstructable, // TODO: Header type fix
    config: {
      placeholder: 'Текст заголовка',
      levels: [1, 2, 3]
    }
  },
  List,
  CheckList,
  Code,
  Marker,
  InlineCode,
  HyperLink
}
