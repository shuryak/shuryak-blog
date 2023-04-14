import {
  FetchRefreshFailure,
  FetchRefreshFailurePayload,
  FetchRefreshRequest, FetchRefreshRequestPayload, FetchRefreshSuccess,
  FetchRefreshSuccessPayload,
  RefreshTypes
} from '../../types/types'

export const fetchRefreshRequest = (payload: FetchRefreshRequestPayload): FetchRefreshRequest => ({
  type: RefreshTypes.FETCH_REFRESH_REQUEST,
  payload
})

export const fetchRefreshSuccess = (payload: FetchRefreshSuccessPayload): FetchRefreshSuccess => ({
  type: RefreshTypes.FETCH_REFRESH_SUCCESS,
  payload
})

export const fetchRefreshFailure = (payload: FetchRefreshFailurePayload) => ({ // TODO: type
  type: RefreshTypes.FETCH_REFRESH_FAILURE,
  payload
})
