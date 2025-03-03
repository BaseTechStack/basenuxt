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
          
          <UFormField label="Description" >
            
            <UInput 
              v-model="formData.description" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="StartDate" >
            
            <UInput 
              v-model="formData.startDate" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="EndDate" >
            
            <UInput 
              v-model="formData.endDate" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="Status" >
            
            <UInput 
              v-model="formData.status" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="Budget" >
            
            <UInput 
              v-model="formData.budget" 
               
              type="number"
            />
            
          </UFormField>
          
          <UFormField label="ClientId" >
            
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
  
  description: '',
  
  startDate: '',
  
  endDate: '',
  
  status: '',
  
  budget: 0,
  
  clientId: 0,
  
})

function closeModal() {
  isOpen.value = false
  
  // Reset form data
  
  formData.name = ''
  
  formData.description = ''
  
  formData.startDate = ''
  
  formData.endDate = ''
  
  formData.status = ''
  
  formData.budget = 0
  
  formData.clientId = 0
  
}

function handleSubmit() {
  emit('submit', { ...formData })
  closeModal()
}
</script>
