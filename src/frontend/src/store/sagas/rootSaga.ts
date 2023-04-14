import {all, fork} from 'redux-saga/effects'
import {loginSaga, refreshSaga} from './loginSaga/loginSaga'

export function* rootSaga() {
  yield all([fork(loginSaga), fork(refreshSaga)])
}
