import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'
import {Provider} from 'react-redux'
import store from './store/store'

const root = createRoot(document.getElementById('app') as HTMLElement)
root.render(
  // <React.StrictMode> // TODO: React.StrictMode + EditorJS (https://github.com/Jungwoo-An/react-editor-js/issues/228)
    <Provider store={store}>
      <App/>
    </Provider>
  // </React.StrictMode>
)
