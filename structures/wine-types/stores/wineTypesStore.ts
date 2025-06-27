//pinia store for WineTypes

// Add TypeScript declaration for import.meta.client
declare global {
  interface ImportMeta {
    client: boolean
  }
}

import type { WineType } from './wineType'
import { WineTypeModel } from './wineType'
import { defineStore } from 'pinia'
import { defineBaseStore, type BaseItem } from '@@/app/stores/baseStore'
import { useWineTypeService } from '../services/wineTypeService'

const VIEW_MODE_KEY = 'winetypes_view_mode'

// Create the WineTypes store using defineStore directly
export const useWineTypesStore = defineStore('winetypes', {
  state: () => ({
    items: [] as WineType[],
    loading: false,
    error: null as string | null,
    viewMode: (typeof window !== 'undefined' ? localStorage.getItem(VIEW_MODE_KEY) as 'grid' | 'table' : null) || 'grid',
    pagination: {
      total: 0,
      page: 1,
      pageSize: 12, // Changed from 10 to 12 to be consistent with our pagination options
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
      console.log('winetypesStore fetch called with page:', currentPage, 'pageSize:', pageSize)
      this.loading = true
      try {
        const winetypeService = useWineTypeService()
        const { items, pagination } = await winetypeService.fetch(currentPage, pageSize || this.pagination.pageSize)
        this.items = items
        this.pagination = { ...this.pagination, ...pagination }
        console.log('Updated pagination state:', this.pagination)
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    },
    
    async create(data: Omit<WineType, 'id'>): Promise<WineType> {
      this.loading = true
      try {
        const winetypeService = useWineTypeService()
        // Convert data to snake_case for API if needed
        const apiData = WineTypeModel.toJson(data, true)
        const { item } = await winetypeService.create(apiData as any)
        this.items.push(item)
        return item
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async update(id: number, data: Partial<Omit<WineType, 'id'>>): Promise<WineType> {
      this.loading = true
      try {
        const winetypeService = useWineTypeService()
        // Convert data to snake_case for API if needed
        const apiData = WineTypeModel.toJson(data, true)
        const { item } = await winetypeService.update(id, apiData as any)
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
        const winetypeService = useWineTypeService()
        await winetypeService.delete(id)
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
export function fetchWineTypes(page = 1, pageSize?: number): Promise<void> {
  console.log('fetchWineTypes called with page:', page, 'pageSize:', pageSize)
  const store = useWineTypesStore()
  return store.fetch(page, pageSize)
}

export function createWineType(winetypeData: Omit<WineType, 'id'>): Promise<WineType> {
  const store = useWineTypesStore()
  return store.create(winetypeData)
}

export function updateWineTypeById(id: number, winetypeData: Partial<Omit<WineType, 'id'>>): Promise<WineType> {
  const store = useWineTypesStore()
  return store.update(id, winetypeData)
}

export function deleteWineType(id: number): Promise<void> {
  const store = useWineTypesStore()
  return store.delete(id)
}
