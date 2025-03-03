<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
    <GridCard 
      v-for="entity in entities" 
      :key="entity.id" 
      :item="entity"
      title="id"
      subtitle="name"
      :fields="[
        
        
        
        
        
        { key: 'email', label: 'Email' },
        
        { key: 'age', label: 'Age' },
        
        { key: 'active', label: 'Active' },
        
      ]"
      timestamp="createdAt"
      hover
    >
      <template #actions>
        <BaseCrudActions 
          structure="entities"
          :item="entity" 
          actions="view,edit,delete"
          @view="$emit('view', $event)"
          @edit="$emit('edit', $event)"
          @delete="$emit('delete', $event)"
        />
      </template>
      
      <template #footer-actions>
        <UButton
          size="xs"
          color="primary"
          variant="ghost"
          icon="i-heroicons-eye"
          @click="$emit('view', entity)"
        >
          View Details
        </UButton>
      </template>
    </GridCard>
  </div>
</template>

<script setup lang="ts">
import type { entity } from '../stores/entity'
import { computed } from 'vue'

const props = defineProps<{
  entity: entity[]
}>()

const entities = computed(() => props.entity)

defineEmits<{
  (e: 'view', entity: entity): void
  (e: 'edit', entity: entity): void
  (e: 'delete', entity: entity): void
}>()
</script>
