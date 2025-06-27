<template>
  <USlideover
    v-model:open="isOpen"
    title="View WineType"
    description="View the details of this WineType."
  >

    <template #body>
      <div class="space-y-6 p-4">
        <div v-for="field in fields" :key="field.label" class="space-y-1">
          <div class="text-sm font-medium text-gray-500 dark:text-gray-400">{{ field.label }}</div>
          <div>{{ field.value }}</div>
        </div>
        
        <div v-if="item" class="space-y-1">
          <div class="text-sm font-medium text-gray-500 dark:text-gray-400">Created</div>
          <div>{{ formatDate(item.createdAt) }}</div>
        </div>
        
        <div v-if="item" class="space-y-1">
          <div class="text-sm font-medium text-gray-500 dark:text-gray-400">Updated</div>
          <div>{{ formatDate(item.updatedAt) }}</div>
        </div>
      </div>
    </template>
    
    <template #footer>
      <div class="flex justify-end p-4">
        <UButton
          variant="soft"
          color="neutral"
          @click="isOpen = false"
        >
          Close
        </UButton>
      </div>
    </template>
  </USlideover>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { format } from 'date-fns'
import { useWineTypesStore } from '../stores/winetypesStore'
import type { WineType } from '../stores/winetype'
import { WineTypeModel } from '../stores/winetype'

const formatDate = (dateString?: string) => {
  if (!dateString) return 'N/A'
  try {
    return format(new Date(dateString), 'PPP')
  } catch (e) {
    return dateString
  }
}

const winetypesStore = useWineTypesStore()

const props = defineProps<{
  open?: boolean
  winetype?: WineType
  title?: string
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
}>()

// Create a computed property for v-model binding with USlideover
const isOpen = computed({
  get() {
    return props.open
  },
  set(value) {
    emit('update:open', value)
  }
})

// No need for a computed title as we're using a hardcoded one

const item = ref<WineType | null>(null)
const fields = ref<{label: string, value: any}[]>([])

const populateFields = () => {
  if (!item.value) return
  
  fields.value = [
    
    { 
      label: 'Name', 
      value: item.value.name
    },
    
    { 
      label: 'Description', 
      value: item.value.description
    },
    
    { 
      label: 'Category', 
      value: item.value.category
    },
    
  ]
}

const fetchData = () => {
  if (props.winetype) {
    // Convert the raw data to a proper model instance if needed
    if (props.winetype.created_at && !props.winetype.createdAt) {
      item.value = WineTypeModel.fromJson(props.winetype as any)
    } else {
      item.value = props.winetype
    }
    populateFields()
  }
}

watch(isOpen, (newVal) => {
  if (newVal && props.winetype) {
    fetchData()
  } else {
    item.value = null
    fields.value = []
  }
})

watch(() => props.winetype, (newVal) => {
  if (isOpen.value && newVal) {
    fetchData()
  }
})
</script>
