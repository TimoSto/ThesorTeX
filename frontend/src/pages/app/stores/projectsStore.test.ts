import { setActivePinia, createPinia } from 'pinia'
import {useProjectsStore} from "@/pages/app/stores/projectsStore";

describe('projects store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  });
  describe('setting data', () => {
    it('set data on empty', () => {
      const store = useProjectsStore();
      store.setProjects([
        {
          Name: "test1",
          Created: '',
          LastModified: '',
          NumberOfEntries: 0
        }
      ]);
      expect(store.projects).toHaveLength(1)
      expect(store.projects[0]).toEqual({
        Name: "test1",
        Created: '',
        LastModified: '',
        NumberOfEntries: 0
      })
    })
    it('rm project', () => {
      const store = useProjectsStore();
      store.projects = [
        {
          Name: "test1",
          Created: '',
          LastModified: '',
          NumberOfEntries: 0
        },
        {
          Name: "test2",
          Created: '',
          LastModified: '',
          NumberOfEntries: 0
        }
      ];
      store.rmProject('test2')
      expect(store.projects).toHaveLength(1)
      expect(store.projects[0]).toEqual({
        Name: "test1",
        Created: '',
        LastModified: '',
        NumberOfEntries: 0
      })
    })
  })
})