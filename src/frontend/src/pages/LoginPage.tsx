import React, {useState, useEffect} from 'react'
import {useDispatch, useSelector} from 'react-redux'
import {fetchLoginRequest} from '../store/actions/userActions/loginActions'
import {RootState} from '../store/reducers/rootReducer'
import { useNavigate } from 'react-router-dom';
import '@styles/page.scss'

export const LoginPage: React.FC = () => {
  const [showError, setShowError] = useState<boolean>(false)
  const [username, setUsername] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const navigate = useNavigate()
  const dispatch = useDispatch()

  const {pending, accessToken, error} = useSelector(
    (state: RootState) => state.login
  )

  function login() {
    dispatch(fetchLoginRequest({
      username: username,
      password: password
    }))
  }

  useEffect(() => {
    if (accessToken) {
      localStorage.setItem('username', username)
      navigate('/')
    }
  }, [accessToken])

  useEffect(() => {
    setShowError(false)
  }, [username, password])

  useEffect(() => {
    if (error != null) {
      setShowError(true)
    }
  }, [pending])

  return (
    <>
      <h1>👤 Авторизация</h1>

      <p className="error-message" style={showError ? {opacity: 100} : {opacity: 0}}>
        🚫 Неверное имя пользователя или пароль.
      </p>

      <form className="account-form">
        <div className="input-block">
          <label htmlFor="login">Имя пользователя:</label>
          <input
            className="account-form-input user-input"
            type="text"
            id="login"
            disabled={pending}
            value={username}
            onChange={e => setUsername(e.target.value)}
          />
        </div>

        <div className="input-block">
          <label htmlFor="pword">Пароль:</label>
          <input
            className="account-form-input user-input"
            type="password"
            id="pword"
            disabled={pending}
            value={password}
            onChange={e => setPassword(e.target.value)}
          />
        </div>

        <input
          className="default-btn user-input"
          type="submit"
          value="Войти"
          disabled={pending}
          onClick={() => login()}
        />
      </form>
    </>
  )
}
