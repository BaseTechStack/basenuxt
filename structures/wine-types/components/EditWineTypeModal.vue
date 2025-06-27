<template>
  <UModal 
    v-model:open="isOpen" 
    :title="title || 'Edit WineType'"
    :description="description || 'Update this winetype'"
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
          
          <UFormField label="Category" >
            
            <UInput 
              v-model="formData.category" 
              
              
                
                
                
                
              
              
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
import type { WineType } from '../stores/winetype'

const props = defineProps<{
  open?: boolean
  winetype?: WineType
  title?: string
  description?: string
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'submit', id: number, data: Partial<WineType>): void
}>()

const isOpen = computed({
  get() {
    return props.open
  },
  set(value) {
    emit('update:open', value)
  }
})

const formData = reactive<Partial<WineType>>({
  
  name: '',
  
  description: '',
  
  category: '',
  
})

// Watch for changes to the winetype prop and update form data
watch(() => props.winetype, (newWineType) => {
  if (newWineType) {
    
    formData.name = newWineType.name || ''
    
    formData.description = newWineType.description || ''
    
    formData.category = newWineType.category || ''
    
  }
}, { immediate: true })

function closeModal() {
  isOpen.value = false
}

function handleSubmit() {
  if (props.winetype?.id) {
    emit('submit', props.winetype.id, { ...formData })
  }
  closeModal()
}
</script>
