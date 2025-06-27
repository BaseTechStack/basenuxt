import type { WineType } from '../stores/wineType'
import type { BasePagination } from '@@/app/types/base'
import { useApi } from '@@/app/composables/useApi'

export const useWineTypes = () => {
  const api = useApi()

  const fetchWineTypes = async (page = 1, pageSize = 10) => {
    console.log('useWineTypes.fetchWineTypes called with page:', page, 'pageSize:', pageSize)
    const result = await api.get(`/wine-types?page=${page}&limit=${pageSize}`)
    console.log('useWineTypes API response:', result)
    const total = result.pagination?.total || 0
    const pagination: BasePagination = {
      total: total,
      page: result.pagination?.page || 1,
      pageSize: pageSize,
      totalPages: Math.max(1, Math.ceil(total / pageSize))
    }
    console.log('useWineTypes returning pagination:', pagination)
    return { winetypes: result.data, pagination }
  }

  const fetchWineTypeById = async (id: number) => {
    const result = await api.get(`/wine-types/${id}`)
    return { winetype: result.data }
  }

  const createWineType = async (winetypeData: Omit<WineType, 'id'>) => {
    const result = await api.post('/wine-types', winetypeData)
    return { winetype: result.data }
  }

  const updateWineType = async (id: number, winetypeData: Partial<Omit<WineType, 'id'>>) => {
    const result = await api.put(`/wine-types/${id}`, winetypeData)
    return { winetype: result.data }
  }

  const deleteWineType = async (id: number) => {
    await api.delete(`/wine-types/${id}`)
  }

  return {
    fetchWineTypes,
    fetchWineTypeById,
    createWineType,
    updateWineType,
    deleteWineType
  }
}
