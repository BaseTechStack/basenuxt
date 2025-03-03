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
          
          <UFormField label="Description" >
            
            <UInput 
              v-model="formData.description" 
              
              
                
                
                
                
              
              
            />
            
          </UFormField>
          
          <UFormField label="Startdate" >
            
            <UInput 
              v-model="formData.startDate" 
              
              
                
                
                
                
              
              
            />
            
          </UFormField>
          
          <UFormField label="Enddate" >
            
            <UInput 
              v-model="formData.endDate" 
              
              
                
                
                
                
              
              
            />
            
          </UFormField>
          
          <UFormField label="Status" >
            
            <UInput 
              v-model="formData.status" 
              
              
                
                
                
                
              
              
            />
            
          </UFormField>
          
          <UFormField label="Budget" >
            
            <UInput 
              v-model="formData.budget" 
              
              
              type="number"
            />
            
          </UFormField>
          
          <UFormField label="Clientid" >
            
            <UInput 
              v-model="formData.clientId" 
              
              
              type="number"
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
  
  description: '',
  
  startDate: '',
  
  endDate: '',
  
  status: '',
  
  budget: 0,
  
  clientId: 0,
  
})

// Watch for changes to the project prop and update form data
watch(() => props.project, (newProject) => {
  if (newProject) {
    
    formData.name = newProject.name || ''
    
    formData.description = newProject.description || ''
    
    formData.startDate = newProject.startDate || ''
    
    formData.endDate = newProject.endDate || ''
    
    formData.status = newProject.status || ''
    
    formData.budget = newProject.budget || 0
    
    formData.clientId = newProject.clientId || 0
    
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
