import type { entity } from '../stores/entity'
import type { BasePagination } from '@@/app/types/base'
import { useApi } from '@@/app/composables/useApi'

interface ApiResponse<T> {
  data: T
  pagination?: {
    total: number
    page: number
    total_pages: number
  }
}

export const useentities = () => {
  const api = useApi()

  const fetchentities = async (page = 1, pageSize = 10) => {
    console.log('useentities.fetchentities called with page:', page, 'pageSize:', pageSize)
    const result = await api.get(`/entities?page=${page}&limit=${pageSize}`)
    console.log('useentities API response:', result)
    const total = result.pagination?.total || 0
    const pagination: BasePagination = {
      total: total,
      page: result.pagination?.page || 1,
      pageSize: pageSize,
      totalPages: Math.max(1, Math.ceil(total / pageSize))
    }
    console.log('useentities returning pagination:', pagination)
    return { entities: result.data, pagination }
  }

  const createentity = async (entityData: Omit<entity, 'id'>) => {
    const result = await api.post('/entities', entityData)
    return { entity: result.data }
  }

  const updateentity = async (id: number, entityData: Partial<Omit<entity, 'id'>>) => {
    const result = await api.put(`/entities/${id}`, entityData)
    return { entity: result.data }
  }

  const deleteentity = async (id: number) => {
    await api.delete(`/entities/${id}`)
  }

  return {
    fetchentities,
    createentity,
    updateentity,
    deleteentity
  }
}
