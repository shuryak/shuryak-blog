import {LoginRequest} from '../../models/token'

// Login

export enum LoginTypes {
  FETCH_LOGIN_REQUEST = "FETCH_LOGIN_REQUEST",
  FETCH_LOGIN_SUCCESS = "FETCH_LOGIN_SUCCESS",
  FETCH_LOGIN_FAILURE = "FETCH_LOGIN_FAILURE"
}

export interface TokenState {
  pending: boolean
  username: string | null,
  accessToken: string | null
  error: string | null
}

export interface FetchLoginRequestPayload {
  username: string
  password: string
}

export interface FetchLoginSuccessPayload {
  accessToken: string
}

export interface FetchLoginFailurePayload {
  error: string
}

export interface FetchLoginRequest {
  type: LoginTypes.FETCH_LOGIN_REQUEST
  payload: FetchLoginRequestPayload
}

export type FetchLoginSuccess = {
  type: LoginTypes.FETCH_LOGIN_SUCCESS
  payload: FetchLoginSuccessPayload
}

export type FetchLoginFailure = {
  type: LoginTypes.FETCH_LOGIN_FAILURE
  payload: FetchLoginFailurePayload
  e: string
}

export type LoginActions =
  | FetchLoginRequest
  | FetchLoginSuccess
  | FetchLoginFailure

// Refresh

export enum RefreshTypes {
  FETCH_REFRESH_REQUEST = "FETCH_REFRESH_REQUEST",
  FETCH_REFRESH_SUCCESS = "FETCH_REFRESH_SUCCESS",
  FETCH_REFRESH_FAILURE = "FETCH_REFRESH_FAILURE"
}

export interface FetchRefreshRequestPayload {
  username: string
}

export interface FetchRefreshSuccessPayload {
  accessToken: string
}

export interface FetchRefreshFailurePayload {
  error: string
}

export interface FetchRefreshRequest {
  type: RefreshTypes.FETCH_REFRESH_REQUEST
  payload: FetchRefreshRequestPayload
}

export type FetchRefreshSuccess = {
  type: RefreshTypes.FETCH_REFRESH_SUCCESS
  payload: FetchRefreshSuccessPayload
}

export type FetchRefreshFailure = {
  type: RefreshTypes.FETCH_REFRESH_FAILURE
  payload: FetchRefreshFailurePayload
  e: string
}

export type RefreshActions =
  | FetchRefreshRequest
  | FetchRefreshSuccess
  | FetchRefreshFailure

// LocalStorage

export enum UsernameTypes {
  FETCH_USERNAME = "FETCH_USERNAME"
}

export interface FetchUsernamePayload {
  username: string | null
}

export interface FetchUsername {
  type: UsernameTypes.FETCH_USERNAME
  payload: FetchUsernamePayload
}

export type UsernameActions =
  | FetchUsername
