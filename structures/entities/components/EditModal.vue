<template>
  <UModal 
    v-model:open="isOpen" 
    :title="title || 'Edit entity'"
    :description="description || 'Update this entity'"
  >
    <template #body>
      <form @submit.prevent="handleSubmit">
        <div class="space-y-4">
          
          <UFormField label="Test" >
            
            <UInput 
              v-model="formData.test" 
              
              
                
                
                
                
              
              
            />
            
          </UFormField>
          
          <UFormField label="Name" >
            
            <UInput 
              v-model="formData.name" 
              
              
                
                
                
                
              
              
            />
            
          </UFormField>
          
          <UFormField label="Email" >
            
            <UInput 
              v-model="formData.email" 
              
              
                type="email"
                
                
                
              
              
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
          Update
        </UButton>
      </div>
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { ref, computed, reactive, watch } from 'vue'
import type { entity } from '../stores/entity'

const props = defineProps<{
  open?: boolean
  entity?: entity
  title?: string
  description?: string
  loading?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'submit', id: number, data: Partial<entity>): void
}>()

const isOpen = computed({
  get() {
    return props.open
  },
  set(value) {
    emit('update:open', value)
  }
})

const formData = reactive<Partial<entity>>({
  
  test: '',
  
  name: '',
  
  email: '',
  
  age: 0,
  
  active: false,
  
})

// Watch for changes to the entity prop and update form data
watch(() => props.entity, (newentity) => {
  if (newentity) {
    
    formData.test = newentity.test || ''
    
    formData.name = newentity.name || ''
    
    formData.email = newentity.email || ''
    
    formData.age = newentity.age || 0
    
    formData.active = newentity.active || false
    
  }
}, { immediate: true })

function closeModal() {
  isOpen.value = false
}

function handleSubmit() {
  if (props.entity?.id) {
    emit('submit', props.entity.id, { ...formData })
  }
  closeModal()
}
</script>
