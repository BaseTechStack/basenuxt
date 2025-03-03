import type { Project } from '../stores/project'
import type { BaseService } from '@@/app/services/baseService'
import { useProjects } from '../composables/useProjects'
import type { BasePagination } from '@@/app/types/base'

// Use a composable pattern to ensure Pinia is initialized before using the store
export const useProjectService = (): BaseService<Project> => {
  // Only create useProjects() when the composable is called, not at module level
  const projectsApi = useProjects()
  
  return {
    async fetch(page = 1, pageSize = 10): Promise<{ items: Project[], pagination: BasePagination }> {
      console.log('ProjectService.fetch called with page:', page, 'pageSize:', pageSize)
      const result = await projectsApi.fetchProjects(page, pageSize)
      console.log('ProjectService API response:', result)
      return {
        items: result.projects.map((project: Project) => ({
          ...project,
          id: Number(project.id)
        })),
        pagination: result.pagination
      }
    },

    async create(data: Omit<Project, 'id'>): Promise<{ item: Project }> {
      const result = await projectsApi.createProject(data)
      return {
        item: {
          ...result.project,
          id: Number(result.project.id)
        }
      }
    },

    async update(id: number, data: Partial<Omit<Project, 'id'>>): Promise<{ item: Project }> {
      const result = await projectsApi.updateProject(id, data)
      return {
        item: {
          ...result.project,
          id: Number(result.project.id)
        }
      }
    },

    async delete(id: number): Promise<void> {
      await projectsApi.deleteProject(id)
    }
  }
}
