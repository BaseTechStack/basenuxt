<template>
  <UModal 
    v-model:open="isOpen" 
    :title="title || 'Edit Project'"
    :description="description || 'Update this project'"
  >
    <template #body>
      <form @submit.prevent="handleSubmit">
        <div class="space-y-4">
          
          <UFormField label="Name" >
            
            <UInput 
              v-model="formData.name" 
              
              
              
            />
            
          </UFormField>
          
          <UFormField label="/users/flakerimismani/base/basenuxt/testproject" >
            
            <UInput 
              v-model="formData./users/flakerimismani/base/basenuxt/testProject" 
              
              
                
                
                
                
              
              
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
          Update
        </UButton>
      </div>
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { ref, computed, reactive, watch } from 'vue'
import type { Project } from '../stores/project'

const props = defineProps<{
  open?: boolean
  project?: Project
  title?: string
  description?: string
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'submit', id: number, data: Partial<Project>): void
}>()

const isOpen = computed({
  get() {
    return props.open
  },
  set(value) {
    emit('update:open', value)
  }
})

const formData = reactive<Partial<Project>>({
  
  name: '',
  
  /users/flakerimismani/base/basenuxt/testProject: '',
  
})

// Watch for changes to the project prop and update form data
watch(() => props.project, (newProject) => {
  if (newProject) {
    
    formData.name = newProject.name || ''
    
    formData./users/flakerimismani/base/basenuxt/testProject = newProject./users/flakerimismani/base/basenuxt/testProject || ''
    
  }
}, { immediate: true })

function closeModal() {
  isOpen.value = false
}

function handleSubmit() {
  if (props.project?.id) {
    emit('submit', props.project.id, { ...formData })
  }
  closeModal()
}
</script>
