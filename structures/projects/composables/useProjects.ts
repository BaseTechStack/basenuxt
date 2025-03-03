import type { Project } from '../stores/project'
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

export const useProjects = () => {
  const api = useApi()

  const fetchProjects = async (page = 1, pageSize = 10) => {
    console.log('useProjects.fetchProjects called with page:', page, 'pageSize:', pageSize)
    const result = await api.get(`/projects?page=${page}&limit=${pageSize}`)
    console.log('useProjects API response:', result)
    const total = result.pagination?.total || 0
    const pagination: BasePagination = {
      total: total,
      page: result.pagination?.page || 1,
      pageSize: pageSize,
      totalPages: Math.max(1, Math.ceil(total / pageSize))
    }
    console.log('useProjects returning pagination:', pagination)
    return { projects: result.data, pagination }
  }

  const createProject = async (projectData: Omit<Project, 'id'>) => {
    const result = await api.post('/projects', projectData)
    return { project: result.data }
  }

  const updateProject = async (id: number, projectData: Partial<Omit<Project, 'id'>>) => {
    const result = await api.put(`/projects/${id}`, projectData)
    return { project: result.data }
  }

  const deleteProject = async (id: number) => {
    await api.delete(`/projects/${id}`)
  }

  return {
    fetchProjects,
    createProject,
    updateProject,
    deleteProject
  }
}
