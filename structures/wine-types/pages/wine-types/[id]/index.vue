<template>
  <div class="wine-type-detail-container">
    <header class="wine-type-detail-header p-4 flex justify-between items-center mb-6">
      <div class="flex items-center">
        <NuxtLink :to="`/wine-types`" class="back-link">
          <UButton
            color="neutral"
            variant="ghost"
            icon="i-lucide-arrow-left"
          >
            Back to WineTypes
          </UButton>
        </NuxtLink>
        <h1 class="text-2xl font-bold ml-4">{{ wineType ? wineType.name : 'Loading WineType...' }}</h1>
      </div>
      
      <div v-if="{{ wineType }}" class="flex items-center space-x-2">
        <UButton
          color="primary"
          icon="i-heroicons-pencil-square"
          @click="isEditing = true"
          v-if="!isEditing"
        >
          Edit
        </UButton>
        <UButton
          color="red"
          icon="i-heroicons-trash"
          @click="showDeleteConfirm = true"
          v-if="!isEditing"
        >
          Delete
        </UButton>
        <UButton
          color="gray"
          @click="isEditing = false"
          v-if="isEditing"
        >
          Cancel
        </UButton>
        <UButton
          color="primary"
          icon="i-heroicons-check"
          @click="saveChanges"
          v-if="{{ isEditing }}"
          :loading="{{ saving }}"
        >
          Save
        </UButton>
      </div>
    </header>

    <div v-if="{{ loading }}" class="flex flex-col items-center justify-center p-12">
      <ULoading />
      <p class="mt-4 text-gray-500">Loading WineType...</p>
    </div>

    <div v-else-if="{{ error }}" class="p-4">
      <UAlert
        title="Error"
        color="error"
        variant="soft"
        icon="i-lucide-alert-triangle"
      >
        <template #description>
          {{ error }}
        </template>
      </UAlert>
    </div>

    <div v-else-if="{{ wineType }}" class="p-4">
      <UCard v-if="{{ !isEditing }}" class="wine-type-detail">
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold">WineType Details</h3>
            <div class="wine-type-meta flex space-x-2">
              <UBadge color="neutral" variant="soft" class="text-xs">
                ID: {{ wineType.id }}
              </UBadge>
              <UBadge color="primary" variant="soft" class="text-xs">
                Created: {{ formatDate(wineType.createdAt) }}
              </UBadge>
              <UBadge color="primary" variant="soft" class="text-xs">
                Updated: {{ formatDate(wineType.updatedAt) }}
              </UBadge>
            </div>
          </div>
        </template>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 py-4">
          
          <UCard class="field-card" :ui="{ body: { padding: 'p-3' } }">
            <div class="flex flex-col">
              <div class="text-sm font-medium text-gray-500">Name</div>
              <div class="mt-1 font-medium">{{ wineType.name || '-' }}</div>
            </div>
          </UCard>
          
          <UCard class="field-card" :ui="{ body: { padding: 'p-3' } }">
            <div class="flex flex-col">
              <div class="text-sm font-medium text-gray-500">Description</div>
              <div class="mt-1 font-medium">{{ wineType.description || '-' }}</div>
            </div>
          </UCard>
          
          <UCard class="field-card" :ui="{ body: { padding: 'p-3' } }">
            <div class="flex flex-col">
              <div class="text-sm font-medium text-gray-500">Category</div>
              <div class="mt-1 font-medium">{{ wineType.category || '-' }}</div>
            </div>
          </UCard>
          
        </div>
      </UCard>

      <FormWineType 
        v-else 
        :wineType="{{ wineType }}" 
        :loading="{{ saving }}"
        @submit="handleUpdate"
      />
    </div>

    <!-- Delete Confirmation Dialog -->
    <UModal v-model="{{ showDeleteConfirm }}">
      <UCard :ui="{ divide: 'divide-y divide-gray-100 dark:divide-gray-800' }">
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
              Delete WineType
            </h3>
            <UButton
              color="gray"
              variant="ghost"
              icon="i-heroicons-x-mark-20-solid"
              class="-my-1"
              @click="showDeleteConfirm = false"
            />
          </div>
        </template>

        <div class="py-4">
          <p class="text-sm text-gray-500">
            Are you sure you want to delete this winetype? This action cannot be undone.
          </p>
        </div>

        <template #footer>
          <div class="flex justify-end gap-x-4">
            <UButton
              label="Cancel"
              color="gray"
              variant="solid"
              @click="showDeleteConfirm = false"
            />
            <UButton
              label="Delete"
              color="red"
              :loading="deleting"
              @click="confirmDelete"
            />
          </div>
        </template>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { WineType } from '@@/wine-types/stores/wineType'
import FormWineType from '@@/wine-types/components/FormWineType.vue'
import { useWineTypeService } from '@@/wine-types/services/wineTypeService'

const route = useRoute()
const id = computed(() => String(route.params.id))
const wineTypeService = useWineTypeService()

const wineType = ref<WineType | null>(null)
const loading = ref(true)
const error = ref('')
const isEditing = ref(false)
const saving = ref(false)
const showDeleteConfirm = ref(false)
const deleting = ref(false)

onMounted(async () => {
  try {
    const itemId = route.params.id
    if (!itemId) {
      error.value = 'No WineType ID provided'
      loading.value = false
      return
    }

    wineType.value = await wineTypeService.fetchById(itemId as string)
    loading.value = false
  } catch (e) {
    error.value = `Error loading WineType: ${e.message}`
    loading.value = false
  }
})

function formatDate(dateString: string) {
  try {
    return format(new Date(dateString), 'PPP')
  } catch (e) {
    return dateString
  }
}

async function handleUpdate(data: Omit<WineType, 'id' | 'createdAt' | 'updatedAt'>) {
  if (!wineType.value) return
  
  saving.value = true
  try {
    const updatedItem = await wineTypeService.update(wineType.value.id, data)
    
    if (updatedItem) {
      wineType.value = updatedItem
      isEditing.value = false
    }
  } catch (e) {
    error.value = `Error updating WineType: ${e.message}`
  } finally {
    saving.value = false
  }
}

async function saveChanges() {
  if (!wineType.value) return
  
  saving.value = true
  try {
    await wineTypeService.update(wineType.value)
    isEditing.value = false
    // Refresh data
    wineType.value = await wineTypeService.fetchById(wineType.value.id)
  } catch (e) {
    error.value = `Error updating WineType: ${e.message}`
  } finally {
    saving.value = false
  }
}

async function deleteItem() {
  if (!wineType.value) return
  
  try {
    await wineTypeService.delete(wineType.value.id)
    // Navigate back to listing page
    navigateTo(`/wine-types`)
  } catch (e) {
    error.value = `Error deleting WineType: ${e.message}`
    showDeleteConfirm.value = false
  }
}
</script>

<style scoped>
.winetype-detail-container {
  min-height: 100%;
  display: flex;
  flex-direction: column;
}

.winetype-detail-header {
  border-bottom: 1px solid var(--color-gray-200);
}

@media (prefers-color-scheme: dark) {
  .winetype-detail-header {
    border-bottom: 1px solid var(--color-gray-800);
  }
}
</style>
