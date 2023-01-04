import { setActivePinia, createPinia } from 'pinia'
import {useErrorSuccessStore} from "./errorSuccessStore";

describe('ErrorSuccessStore', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  });

  describe('hanldeResponse', () => {
    it('set error', () => {
      const store = useErrorSuccessStore();
      expect(store.errorMessage).toEqual('');
      expect(store.successMessage).toEqual('');
      store.handleResponse(false, 'test 1', 'error1')
      expect(store.errorMessage).toEqual('error1');
      expect(store.successMessage).toEqual('');
    })
    it('set success', () => {
      const store = useErrorSuccessStore();
      expect(store.errorMessage).toEqual('');
      expect(store.successMessage).toEqual('');
      store.handleResponse(true, 'test 1', 'error1')
      expect(store.errorMessage).toEqual('');
      setTimeout(() => {
        expect(store.successMessage).toEqual('test 1');
      })
    })
    it('clean', () => {
      const store = useErrorSuccessStore();
      store.errorMessage = 'test';
      store.successMessage = 'test2';
      store.clean();
      expect(store.errorMessage).toEqual('');
      expect(store.successMessage).toEqual('');
    })
  })
})
