<template>
  <USlideover
    v-model:open="isOpen"
  >
    <template #header>
      <div class="flex items-center justify-between">
        <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
          View Project
        </h3>
      </div>
    </template>
    <template #content>
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
import { useProjectsStore } from '../stores/projectsStore'
import type { Project } from '../stores/project'

const formatDate = (dateString: string) => {
  try {
    return format(new Date(dateString), 'PPP')
  } catch (e) {
    return dateString
  }
}

const projectsStore = useProjectsStore()

const props = defineProps<{
  open?: boolean
  project?: Project
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

const item = ref<Project | null>(null)
const fields = ref<{label: string, value: any}[]>([])

const populateFields = () => {
  if (!item.value) return
  
  fields.value = [
    
    { label: 'Name', value: item.value.name },
    
    { label: '/users/flakerimismani/base/basenuxt/testProject', value: item.value./users/flakerimismani/base/basenuxt/testProject },
    
  ]
}

const fetchData = () => {
  if (props.project) {
    item.value = props.project
    populateFields()
  }
}

watch(isOpen, (newVal) => {
  if (newVal && props.project) {
    fetchData()
  } else {
    item.value = null
    fields.value = []
  }
})

watch(() => props.project, (newVal) => {
  if (isOpen.value && newVal) {
    fetchData()
  }
})
</script>
