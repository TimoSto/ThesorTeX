import { setActivePinia, createPinia } from 'pinia'
import {useProjectDataStore} from "@/pages/app/stores/projectDataStore";
import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";

describe('projectData store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  });
  describe('setting data', () => {
    it('set data on empty', () => {
      const store = useProjectDataStore();
      store.setData({
        Entries: [
          {
            Key: 'test',
            Fields: ['t', 't2']
          } as BibEntry
        ],
        Categories: [
          {
            Name: 'tc',
          } as BibCategory
        ]
      });
      expect(store.entries).toHaveLength(1)
      expect(store.entries[0]).toEqual({
        Key: 'test',
        Fields: ['t', 't2']
      })
      expect(store.categories).toHaveLength(1)
      expect(store.categories[0]).toEqual({
        Name: 'tc'
      })
    })
  })
})
