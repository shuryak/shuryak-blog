import {FetchUsername, FetchUsernamePayload, UsernameTypes} from '../../types/types'

export const fetchUsername = (payload: FetchUsernamePayload): FetchUsername => ({
  type: UsernameTypes.FETCH_USERNAME,
  payload
})
