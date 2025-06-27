<template>
  <UModal 
    v-model:open="isOpen" 
    :title="title || 'New WineType'"
    :description="description || 'Add a new winetype to your system'"
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
          
          <UFormField label="Category" >
            
            <UInput 
              v-model="formData.category" 
               
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
import type { WineType } from '../stores/winetype'

const props = defineProps<{
  modelValue?: boolean
  title?: string
  description?: string
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'submit', data: Partial<WineType>): void
  (e: 'winetype-added', data: Partial<WineType>): void
}>()

const isOpen = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  }
})

const formData = reactive<Partial<WineType>>({
  
  name: '',
  
  description: '',
  
  category: '',
  
})

function closeModal() {
  isOpen.value = false
  
  // Reset form data
  
  formData.name = ''
  
  formData.description = ''
  
  formData.category = ''
  
}

function handleSubmit() {
  emit('winetype-added', { ...formData })
  closeModal()
}
</script>
