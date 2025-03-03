<template>
  <div>
    <div class="p-4">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-bold">entities</h1>
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
            Add entity
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
      <Grid 
        v-if="store.viewMode === 'grid' && store.items.length > 0" 
        :entity="store.items" 
        @edit="editentity" 
        @delete="deleteentity" 
        @view="viewentity"
        :current-page="store.pagination.page"
        :page-size="store.pagination.pageSize"
      />
      <Table 
        v-else-if="store.viewMode === 'table' && store.items.length > 0"
        :entity="store.items" 
        @edit="editentity" 
        @delete="deleteentity" 
        @view="viewentity"
        :current-page="store.pagination.page"
        :page-size="store.pagination.pageSize"
      />

      <!-- Empty State -->
      <div v-if="store.items.length === 0" class="text-center py-12">
        <h3 class="mt-4 text-lg font-medium text-gray-900">No entities</h3>
        <p class="mt-1 text-sm text-gray-500">Get started by creating a new entity.</p>
        <div class="mt-6">
          <UButton color="primary" @click="modalState.add.isOpen = true">
            Add entity
          </UButton>
        </div>
      </div>
    </div>

    <!-- Modals -->
    <AddModal
      v-model="modalState.add.isOpen" 
      @entity-added="handleentityAdded"
    />
    
    <EditModal
      v-model:open="modalState.edit.isOpen" 
      :entity="modalState.edit.entity"
      :loading="modalState.edit.loading"
      @submit="handleEditSubmit"
    />
    
    <ViewModal
      v-model:open="modalState.view.isOpen" 
      :entity="modalState.view.entity"
    /> 
    
    <DeleteModal
      v-model:open="modalState.delete.isOpen"
      :entity="modalState.delete.entity"
      :loading="modalState.delete.loading"
      @confirm="confirmDelete"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useentitiesStore } from '../../stores/entitiesStore'
import type { entity } from '../../stores/entity'
import { format } from 'date-fns'

interface ModalState {
  add: {
    isOpen: boolean
    loading?: boolean
  }
  edit: {
    isOpen: boolean
    entity?: entity
    loading?: boolean
  }
  view: {
    isOpen: boolean
    entity?: entity
  }
  delete: {
    isOpen: boolean
    entity?: entity
    loading?: boolean
  }
}

const store = useentitiesStore()

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

function editentity(entity: entity) {
  modalState.value.edit.entity = entity
  modalState.value.edit.isOpen = true
}

function viewentity(entity: entity) {
  modalState.value.view.entity = entity
  modalState.value.view.isOpen = true
}

function deleteentity(entity: entity) {
  modalState.value.delete.entity = entity
  modalState.value.delete.isOpen = true
}

function switchViewToEdit() {
  const entity = modalState.value.view.entity
  closeModal('view')
  if (entity) {
    editentity(entity)
  }
}

function closeModal(type: 'add' | 'edit' | 'view') {
  modalState.value[type].isOpen = false
  if (type === 'edit' || type === 'view') {
    modalState.value[type].entity = undefined
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

async function handleentityAdded(data: Omit<entity, 'id'>) {
  modalState.value.add.loading = true
  try {
    await store.create(data)
    modalState.value.add.isOpen = false
  } catch (error) {
    console.error('Error adding entity:', error)
  } finally {
    modalState.value.add.loading = false
  }
}

async function handleEditSubmit(id: number, data: Partial<entity>) {
  modalState.value.edit.loading = true
  try {
    await store.update(id, data)
    modalState.value.edit.isOpen = false
  } catch (error) {
    console.error('Error updating entity:', error)
  } finally {
    modalState.value.edit.loading = false
  }
}

async function confirmDelete() {
  if (!modalState.value.delete.entity) return
  
  modalState.value.delete.loading = true
  try {
    await store.delete(modalState.value.delete.entity.id)
    modalState.value.delete.isOpen = false
  } catch (error) {
    console.error('Error deleting entity:', error)
  } finally {
    modalState.value.delete.loading = false
  }
}
</script>
