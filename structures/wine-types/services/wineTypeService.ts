import type { WineType } from '../stores/wineType'
import { WineTypeModel } from '../stores/wineType'
import type { BaseService } from '@@/app/services/baseService'
import { useWineTypes } from '../composables/useWineTypes'
import type { BasePagination } from '@@/app/types/base'

// Use a composable pattern to ensure Pinia is initialized before using the store
export const useWineTypeService = (): BaseService<WineType> => {
  // Only create useWineTypes() when the composable is called, not at module level
  const winetypesApi = useWineTypes()
  
  return {
    async fetch(page = 1, pageSize = 10): Promise<{ items: WineType[], pagination: BasePagination }> {
      console.log('WineTypeService.fetch called with page:', page, 'pageSize:', pageSize)
      const result = await winetypesApi.fetchWineTypes(page, pageSize)
      console.log('WineTypeService API response:', result)
      return {
        items: WineTypeModel.fromJsonList(result.winetypes),
        pagination: result.pagination
      }
    },
    async fetchById(id: number): Promise<{ item: WineType }> {
      const result = await winetypesApi.fetchWineTypeById(id)
      return {
        item: WineTypeModel.fromJson(result.winetype)
      }
    },
    async create(data: Omit<WineType, 'id'>): Promise<{ item: WineType }> {
      const result = await winetypesApi.createWineType(data)
      return {
        item: WineTypeModel.fromJson(result.winetype)
      }
    },

    async update(id: number, data: Partial<Omit<WineType, 'id'>>): Promise<{ item: WineType }> {
      const result = await winetypesApi.updateWineType(id, data)
      return {
        item: WineTypeModel.fromJson(result.winetype)
      }
    },

    async delete(id: number): Promise<void> {
      await winetypesApi.deleteWineType(id)
    }
  }
}
