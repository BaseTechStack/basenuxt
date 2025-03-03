//pinia store for entities

// Add TypeScript declaration for import.meta.client
declare global {
  interface ImportMeta {
    client: boolean
  }
}

import type { entity } from './entity'
import { defineStore } from 'pinia'
import { defineBaseStore, type BaseItem } from '@@/app/stores/baseStore'
import { useentityService } from '../services/entityService'

const VIEW_MODE_KEY = 'entities_view_mode'

// Create the entities store using defineStore directly
export const useentitiesStore = defineStore('entities', {
  state: () => ({
    items: [] as entity[],
    loading: false,
    error: null as string | null,
    viewMode: (typeof window !== 'undefined' ? localStorage.getItem(VIEW_MODE_KEY) as 'grid' | 'table' : null) || 'grid',
    pagination: {
      total: 0,
      page: 1,
      pageSize: 10,
      totalPages: 0
    }
  }),
  
  getters: {
    getItemById: (state) => (id: number) => {
      return state.items.find(item => item.id === id)
    }
  },
  
  actions: {
    setViewMode(mode: 'grid' | 'table') {
      this.viewMode = mode
      if (typeof window !== 'undefined') {
        localStorage.setItem(VIEW_MODE_KEY, mode)
      }
    },
    
    async fetch(page?: number, pageSize?: number) {
      // Use the provided page or the current page from state
      const currentPage = page !== undefined ? page : this.pagination.page
      console.log('entitiesStore fetch called with page:', currentPage, 'pageSize:', pageSize)
      this.loading = true
      try {
        const entityService = useentityService()
        const { items, pagination } = await entityService.fetch(currentPage, pageSize || this.pagination.pageSize)
        this.items = items
        this.pagination = { ...this.pagination, ...pagination }
        console.log('Updated pagination state:', this.pagination)
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    },
    
    async create(data: Omit<entity, 'id'>): Promise<entity> {
      this.loading = true
      try {
        const entityService = useentityService()
        const { item } = await entityService.create(data)
        this.items.push(item)
        return item
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async update(id: number, data: Partial<Omit<entity, 'id'>>): Promise<entity> {
      this.loading = true
      try {
        const entityService = useentityService()
        const { item } = await entityService.update(id, data)
        const index = this.items.findIndex(i => i.id === id)
        if (index !== -1) {
          this.items[index] = item
        }
        return item
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async delete(id: number): Promise<void> {
      this.loading = true
      try {
        const entityService = useentityService()
        await entityService.delete(id)
        const index = this.items.findIndex(i => i.id === id)
        if (index !== -1) {
          this.items.splice(index, 1)
        }
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    }
  }
})

// Adding specialized methods to the store
// Instead of extending, we'll use the store object directly
// and add our specialized methods that call the base methods

// Add these helper methods to provide more descriptive names
// These functions should only be called from within components or other composables
// where Pinia is properly initialized
export function fetchentities(page = 1, pageSize?: number): Promise<void> {
  console.log('fetchentities called with page:', page, 'pageSize:', pageSize)
  const store = useentitiesStore()
  return store.fetch(page, pageSize)
}

export function createentity(entityData: Omit<entity, 'id'>): Promise<entity> {
  const store = useentitiesStore()
  return store.create(entityData)
}

export function updateentityById(id: number, entityData: Partial<Omit<entity, 'id'>>): Promise<entity> {
  const store = useentitiesStore()
  return store.update(id, entityData)
}

export function deleteentity(id: number): Promise<void> {
  const store = useentitiesStore()
  return store.delete(id)
}
