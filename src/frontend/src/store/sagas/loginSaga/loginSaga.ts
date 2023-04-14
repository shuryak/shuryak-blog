import axios, {AxiosError, AxiosInstance, AxiosResponse, InternalAxiosRequestConfig} from 'axios'
import { snakeCase, camelCase, mapKeys } from 'lodash'
import {all, call, put, takeLatest} from 'redux-saga/effects'
import {
  fetchLoginFailure,
  fetchLoginSuccess
} from '../../actions/userActions/loginActions'
import {
  FetchLoginRequest,
  FetchLoginRequestPayload,
  FetchRefreshRequest,
  FetchRefreshRequestPayload,
  LoginTypes, RefreshTypes
} from '../../types/types'
import {fetchRefreshFailure, fetchRefreshSuccess} from '../../actions/userActions/refreshActions'

const api = axios.create({
  baseURL: 'http://localhost:8080/',
  timeout: 1000,
  withCredentials: true
})
// https://morioh.com/p/8e8b33c25ea1
api.interceptors.response.use((response: AxiosResponse) => {
  console.log('Cookie from:', response)
  if (response.data && response.headers['content-type'].indexOf('application/json') !== 1) {
    response.data = mapKeys(response.data, (value, key) => camelCase(key))
  }
  return response
})
api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  const newConfig: InternalAxiosRequestConfig = {...config}
  // if (newConfig.headers['content-type'].indexOf('multipart/json') !== 1) {
  //   return newConfig
  // }
  if (config.params) {
    newConfig.params = mapKeys(config.params, (value, key) => snakeCase(key))
  }
  if (config.data) {
    newConfig.data = mapKeys(config.data, (value, key) => snakeCase(key))
  }
  return newConfig
})

const getToken = (payload: FetchLoginRequestPayload) => api.post<string>(
  'user/login',
  {
    username: payload.username,
    password: payload.password,
  },
)

const refreshToken = (payload: FetchRefreshRequestPayload) => api.get<string>(
  'user/refresh_session',
  {
    params: {
      username: payload.username
    },
    withCredentials: true
  }
)

export interface ResponseGenerator {
  config?: any,
  data?: any,
  headers?: any,
  request?: any,
  status?: number,
  statusText?: string
} // TODO: ?

function* fetchTokenSaga(payload: FetchLoginRequest) {
  try {
    const response: ResponseGenerator = yield call(getToken, payload.payload)
    console.log(response)
    yield put(
      fetchLoginSuccess({
        accessToken: response.data.accessToken
      })
    )
  } catch (e) {
    if (e instanceof Error) {
      yield put(
        fetchLoginFailure({
          error: e.message
        })
      )
    }
  }
}

function* fetchRefreshTokenSaga(payload: FetchRefreshRequest) {
  try {
    const response: ResponseGenerator = yield call(refreshToken, payload.payload)
    console.log(response)
    yield put(
      fetchRefreshSuccess({
        accessToken: response.data.accessToken
      })
    )
  } catch (e) {
    if (e instanceof Error) {
      yield put(
        fetchRefreshFailure({
          error: e.message
        })
      )
    }
  }
}

export function* loginSaga() {
  yield all([takeLatest(LoginTypes.FETCH_LOGIN_REQUEST, fetchTokenSaga)])
}

export function* refreshSaga() {
  yield all([takeLatest(RefreshTypes.FETCH_REFRESH_REQUEST, fetchRefreshTokenSaga)])
}

// export default loginSaga // todo: default
