import type { entity } from '../stores/entity'
import type { BaseService } from '@@/app/services/baseService'
import { useentities } from '../composables/useentities'
import type { BasePagination } from '@@/app/types/base'

// Use a composable pattern to ensure Pinia is initialized before using the store
export const useentityService = (): BaseService<entity> => {
  // Only create useentities() when the composable is called, not at module level
  const entitiesApi = useentities()
  
  return {
    async fetch(page = 1, pageSize = 10): Promise<{ items: entity[], pagination: BasePagination }> {
      console.log('entityService.fetch called with page:', page, 'pageSize:', pageSize)
      const result = await entitiesApi.fetchentities(page, pageSize)
      console.log('entityService API response:', result)
      return {
        items: result.entities.map((entity: entity) => ({
          ...entity,
          id: Number(entity.id)
        })),
        pagination: result.pagination
      }
    },

    async create(data: Omit<entity, 'id'>): Promise<{ item: entity }> {
      const result = await entitiesApi.createentity(data)
      return {
        item: {
          ...result.entity,
          id: Number(result.entity.id)
        }
      }
    },

    async update(id: number, data: Partial<Omit<entity, 'id'>>): Promise<{ item: entity }> {
      const result = await entitiesApi.updateentity(id, data)
      return {
        item: {
          ...result.entity,
          id: Number(result.entity.id)
        }
      }
    },

    async delete(id: number): Promise<void> {
      await entitiesApi.deleteentity(id)
    }
  }
}
