import {
  LoginActions,
  LoginTypes,
  RefreshActions,
  RefreshTypes,
  TokenState,
  UsernameActions,
  UsernameTypes
} from '../../types/types'

const initialState: TokenState = {
  pending: false,
  username: null,
  accessToken: null,
  error: null
}

export default (state = initialState, action: LoginActions | RefreshActions | UsernameActions) => {
  switch (action.type) {
    case LoginTypes.FETCH_LOGIN_REQUEST:
      return {
        ...state,
        username: action.payload.username,
        pending: true,
      }
    case LoginTypes.FETCH_LOGIN_SUCCESS:
      return {
        ...state,
        pending: false,
        accessToken: action.payload.accessToken,
        error: null
      }
    case LoginTypes.FETCH_LOGIN_FAILURE:
      return {
        ...state,
        pending: false,
        accessToken: null,
        error: action.payload.error
      }
    case RefreshTypes.FETCH_REFRESH_REQUEST:
      return {
        ...state,
        pending: true,
      }
    case RefreshTypes.FETCH_REFRESH_SUCCESS:
      return {
        ...state,
        pending: false,
        accessToken: action.payload.accessToken,
      }
    case RefreshTypes.FETCH_REFRESH_FAILURE:
      return {
        ...state,
        pending: false,
        accessToken: null,
        error: action.payload.error
      }
    case UsernameTypes.FETCH_USERNAME:
      if (action.payload.username === null) {
        return {
          ...state,
          username: null,
          accessToken: null
        }
      }
      return {
        ...state,
        username: action.payload.username,
      }
    default:
      return {...state}
  }
}
