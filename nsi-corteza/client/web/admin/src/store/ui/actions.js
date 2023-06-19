import {LOADER_DEC, LOADER_HIDE, LOADER_INC,} from './types'

export default {
  incLoader ({ commit }) {
    commit(LOADER_INC)
  },

  decLoader ({ commit }) {
    commit(LOADER_DEC)
  },

  hideLoader ({ commit }) {
    commit(LOADER_HIDE)
  },
}
