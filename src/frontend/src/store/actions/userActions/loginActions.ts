import {
  FetchLoginFailurePayload,
  FetchLoginRequest,
  FetchLoginRequestPayload,
  FetchLoginSuccessPayload,
  LoginTypes
} from '../../types/types'

export const fetchLoginRequest = (payload: FetchLoginRequestPayload): FetchLoginRequest => ({
  type: LoginTypes.FETCH_LOGIN_REQUEST,
  payload
})

export const fetchLoginSuccess = (payload: FetchLoginSuccessPayload) => ({
  type: LoginTypes.FETCH_LOGIN_SUCCESS,
  payload
})

export const fetchLoginFailure = (payload: FetchLoginFailurePayload) => ({
  type: LoginTypes.FETCH_LOGIN_FAILURE,
  payload
})
