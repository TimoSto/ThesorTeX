import {createPinia, setActivePinia} from "pinia";
import {useUnsaveCloseStore} from "@/pages/app/stores/unsaveCloseStore";

describe("unsafeCloseStore", function () {
  beforeEach(() => {
    setActivePinia(createPinia())
  })
  describe('triggering and resolving', () => {
    it('trigger and resolve', () => {
      const store = useUnsaveCloseStore();

      store.triggered = true;

      store.trigger(0).then(accepted => {
        store.triggered = false;
      })

      expect(store.triggered).toBe(true);

      store.promiseResolve(true);

      setTimeout(() => {
        expect(store.triggered).toBe(false);
      })
    })
  })
});
