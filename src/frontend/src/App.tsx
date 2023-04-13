import React, {useEffect, useRef} from 'react'
import { createBrowserRouter, RouterProvider, BrowserRouter, HashRouter, Route, Routes } from 'react-router-dom';
import {useDispatch, useSelector} from 'react-redux'
import Header from './components/Header'
import ArticlesPage from './pages/ArticlesPage'
import AboutPage from './pages/AboutPage'
import EditorPage from './pages/EditorPage'
import PageContainer from './components/PageContainer'
import {LoginPage} from './pages/LoginPage'
import '@styles/main.scss'
import {RootState} from './store/reducers/rootReducer'
import {fetchRefreshRequest} from './store/actions/userActions/refreshActions'
import {fetchLoginRequest} from './store/actions/userActions/loginActions'
import {fetchUsername} from './store/actions/userActions/usernameActions'

const App = () => {
  const dispatch = useDispatch()
  const {pending, username, accessToken, error} = useSelector(
    (state: RootState) => state.login
  )

  function getUsername() {
    dispatch(fetchUsername({
      username: localStorage.getItem('username')
    }))
  }

  const calledOnce = useRef(false);

  useEffect(() => {
    if (calledOnce.current) {
      return;
    }

    getUsername()
    calledOnce.current = true;
  }, [])

  return (
    <HashRouter>
      <Routes>
        <Route path="/" element={<Header/>}>
          <Route path="/" element={<div className="wrapper"><PageContainer><ArticlesPage/></PageContainer></div>}/>
          <Route path="/about" element={<div className="wrapper"><PageContainer><AboutPage/></PageContainer></div>}/>
            <Route path="/editor" element={<div className="wrapper"><PageContainer><EditorPage/></PageContainer></div>}/>
            <Route path="/login" element={<div className="wrapper"><PageContainer><LoginPage/></PageContainer></div>}/>
        </Route>
      </Routes>
    </HashRouter>
  )
}

export default App
