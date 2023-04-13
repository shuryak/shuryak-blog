import React, {useEffect} from 'react'
import {useDispatch, useSelector} from 'react-redux'
import { useNavigate } from 'react-router-dom';
import { Link, Outlet } from 'react-router-dom'
import '@styles/header.scss'
import {RootState} from '../store/reducers/rootReducer'
import {fetchUsername} from '../store/actions/userActions/usernameActions'

export const Header: React.FunctionComponent = () => {
  const dispatch = useDispatch()
  const {pending, username, accessToken, error} = useSelector(
    (state: RootState) => state.login
  )
  const navigate = useNavigate()

  function clearUsername() {
    localStorage.removeItem('username')
    dispatch(fetchUsername({
      username: null
    }))
  }

  return (
    <>
      <header>
        <p className="name" onClick={() => navigate('/')}>shuryak.com</p>
        <p className="extra-name">Сделано с <span className="red">❤</span> на микросервисах.</p>
      </header>
      <div className="actions">
        <nav>
          <Link to="/">Последние статьи</Link>
          <Link to="/about">Обо мне</Link>
          {username &&
            <>
              <Link to="/editor">Редактор</Link>
              <a href="#" onClick={() => clearUsername()}>Выйти</a>
            </>
          }
          {!username &&
              <Link to="/login">Авторизация</Link>
          }
        </nav>
      </div>
      <Outlet/>
    </>
  )
}

export default Header
