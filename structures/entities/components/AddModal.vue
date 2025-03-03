<template>
  <UModal 
    v-model:open="isOpen" 
    :title="title || 'New entity'"
    :description="description || 'Add a new entity to your system'"
  >
    <template #body>
      <form @submit.prevent="handleSubmit">
        <div class="space-y-4">
          
          <UFormField label="Test" >
            
            <UInput 
              v-model="formData.test" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="Name" >
            
            <UInput 
              v-model="formData.name" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="Email" >
            
            <UInput 
              v-model="formData.email" 
               
              type="text"
            />
            
          </UFormField>
          
          <UFormField label="Age" >
            
            <UInput 
              v-model="formData.age" 
               
              type="number"
            />
            
          </UFormField>
          
          <UFormField label="Active" >
            
            <UCheckbox v-model="formData.active"  />
            
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
import type { entity } from '../stores/entity'

const props = defineProps<{
  modelValue?: boolean
  title?: string
  description?: string
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'submit', data: Partial<entity>): void
}>()

const isOpen = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  }
})

const formData = reactive<Partial<entity>>({
  
  test: '',
  
  name: '',
  
  email: '',
  
  age: 0,
  
  active: false,
  
})

function closeModal() {
  isOpen.value = false
  
  // Reset form data
  
  formData.test = ''
  
  formData.name = ''
  
  formData.email = ''
  
  formData.age = 0
  
  formData.active = false
  
}

function handleSubmit() {
  emit('submit', { ...formData })
  closeModal()
}
</script>
