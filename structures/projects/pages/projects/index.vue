<template>
  <div>
    <div class="p-4">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-bold">Projects</h1>
        <div class="flex items-center gap-4">
          <div class="flex items-center space-x-2">
            <UButtonGroup size="xl" class="border border-gray-200 dark:border-gray-800 rounded-md">
              <UButton
                :color="store.viewMode === 'grid' ? 'primary' : 'neutral'"
                @click="store.setViewMode('grid')"
                icon="i-heroicons-squares-2x2"
                variant="ghost"
                size="sm"
              />
              <UButton
                :color="store.viewMode === 'table' ? 'primary' : 'neutral'"
                @click="store.setViewMode('table')"
                icon="i-heroicons-table-cells"
                variant="ghost"
                size="sm"
              />
            </UButtonGroup>
          </div>
        
          <UButton
            color="primary"
            icon="i-heroicons-plus"
            @click="modalState.add.isOpen = true"
          >
            Add Project
          </UButton>
        </div>
      </div>
      
      <!-- Pagination -->
      <div class="mt-6 flex py-5 justify-between items-center">
        <BasePerPage
          :pageSize="store.pagination.pageSize"
          @update:pageSize="handlePageSizeChange"
          :options="store.viewMode === 'grid' ? [12, 24, 36, 48] : [10, 20, 50, 100]"
        />
      
        <BasePagination
          :total="store.pagination.total"
          :page="store.pagination.page"
          :items-per-page="store.pagination.pageSize"
          :total-pages="store.pagination.totalPages"
          @update:page="handlePageChange"
        />
      </div>
      <!-- Page Content -->
      <ProjectGrid 
        v-if="store.viewMode === 'grid' && store.items.length > 0" 
        :project="store.items" 
        @edit="editProject" 
        @delete="deleteProject" 
        @view="viewProject"
        :current-page="store.pagination.page"
        :page-size="store.pagination.pageSize"
      />
      <ProjectTable 
        v-else-if="store.viewMode === 'table' && store.items.length > 0"
        :project="store.items" 
        @edit="editProject" 
        @delete="deleteProject" 
        @view="viewProject"
        :current-page="store.pagination.page"
        :page-size="store.pagination.pageSize"
      />

      <!-- Empty State -->
      <div v-if="store.items.length === 0" class="text-center py-12">
        <h3 class="mt-4 text-lg font-medium text-gray-900">No Projects</h3>
        <p class="mt-1 text-sm text-gray-500">Get started by creating a new project.</p>
        <div class="mt-6">
          <UButton color="primary" @click="modalState.add.isOpen = true">
            Add Project
          </UButton>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <AddProjectModal
      v-model="modalState.add.isOpen" 
      @project-added="handleProjectAdded"
    />
    
    <EditProjectModal
      v-model:open="modalState.edit.isOpen" 
      :project="modalState.edit.project"
      :loading="modalState.edit.loading"
      @submit="handleEditSubmit"
    />
    
    <ViewProjectModal
      v-model:open="modalState.view.isOpen" 
      :project="modalState.view.project"
    />
    
    <DeleteProjectModal
      v-model:open="modalState.delete.isOpen" 
      :project="modalState.delete.project"
      :loading="modalState.delete.loading"
      @confirm="confirmDelete"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useProjectsStore } from '../../stores/projectsStore'
import type { Project } from '../../stores/project'
import { format } from 'date-fns'

interface ModalState {
  add: {
    isOpen: boolean
    loading?: boolean
  }
  edit: {
    isOpen: boolean
    project?: Project
    loading?: boolean
  }
  view: {
    isOpen: boolean
    project?: Project
  }
  delete: {
    isOpen: boolean
    project?: Project
    loading?: boolean
  }
}

const store = useProjectsStore()

const modalState = ref<ModalState>({
  add: { isOpen: false },
  edit: { isOpen: false },
  view: { isOpen: false },
  delete: { isOpen: false }
})


onMounted(async () => {
  await store.fetch()
})

function formatDate(dateString: string) {
  try {
    return format(new Date(dateString), 'PPP')
  } catch (e) {
    return dateString
  }
}

function editProject(project: Project) {
  modalState.value.edit.project = project
  modalState.value.edit.isOpen = true
}

function viewProject(project: Project) {
  modalState.value.view.project = project
  modalState.value.view.isOpen = true
}

function deleteProject(project: Project) {
  modalState.value.delete.project = project
  modalState.value.delete.isOpen = true
}

function switchViewToEdit() {
  const project = modalState.value.view.project
  closeModal('view')
  if (project) {
    editProject(project)
  }
}

function closeModal(type: 'add' | 'edit' | 'view') {
  modalState.value[type].isOpen = false
  if (type === 'edit' || type === 'view') {
    modalState.value[type].project = undefined
  }
}

function handlePageChange(page: number) {
  console.log('handlePageChange called with page:', page)
  store.pagination.page = page
  store.fetch(page) // Explicitly pass the page parameter
}

function handlePageSizeChange(size: number) {
  store.pagination.pageSize = size
  store.pagination.page = 1
  store.fetch()
}

async function handleProjectAdded(data: Omit<Project, 'id'>) {
  modalState.value.add.loading = true
  try {
    await store.create(data)
    modalState.value.add.isOpen = false
  } catch (error) {
    console.error('Error adding project:', error)
  } finally {
    modalState.value.add.loading = false
  }
}

async function handleEditSubmit(id: number, data: Partial<Project>) {
  modalState.value.edit.loading = true
  try {
    await store.update(id, data)
    modalState.value.edit.isOpen = false
  } catch (error) {
    console.error('Error updating project:', error)
  } finally {
    modalState.value.edit.loading = false
  }
}

async function confirmDelete() {
  if (!modalState.value.delete.project) return
  
  modalState.value.delete.loading = true
  try {
    await store.delete(modalState.value.delete.project.id)
    modalState.value.delete.isOpen = false
  } catch (error) {
    console.error('Error deleting project:', error)
  } finally {
    modalState.value.delete.loading = false
  }
}
</script>
