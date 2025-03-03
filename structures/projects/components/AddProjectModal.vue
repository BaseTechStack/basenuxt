<template>
  <UModal 
    v-model:open="isOpen" 
    :title="title || 'New Project'"
    :description="description || 'Add a new project to your system'"
  >
    <template #body>
      <form @submit.prevent="handleSubmit">
        <div class="space-y-4">
          
          <UFormField label="Name" >
            
            <UInput 
              v-model="formData.name" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="/users/flakerimismani/base/basenuxt/testProject" >
            
            <UInput 
              v-model="formData./users/flakerimismani/base/basenuxt/testProject" 
               
              type="text"
            />
            
          </UFormField>
          
        </div>
      </form>
    </template>

    <template #footer>
      <div class="flex justify-end gap-2">
        <UButton
          color="neutral"
          variant="ghost"
          @click="closeModal"
        >
          Cancel
        </UButton>
        <UButton
          color="primary"
          :loading="loading"
          @click="handleSubmit"
        >
          Save
        </UButton>
      </div>
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { ref, computed, reactive } from 'vue'
import type { Project } from '../stores/project'

const props = defineProps<{
  modelValue?: boolean
  title?: string
  description?: string
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'submit', data: Partial<Project>): void
}>()

const isOpen = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  }
})

const formData = reactive<Partial<Project>>({
  
  name: '',
  
  /users/flakerimismani/base/basenuxt/testProject: '',
  
})

function closeModal() {
  isOpen.value = false
  
  // Reset form data
  
  formData.name = ''
  
  formData./users/flakerimismani/base/basenuxt/testProject = ''
  
}

function handleSubmit() {
  emit('submit', { ...formData })
  closeModal()
}
</script>
