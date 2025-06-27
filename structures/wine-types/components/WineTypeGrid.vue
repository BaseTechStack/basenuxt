<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
    <WineTypeGridCard 
      v-for="winetype in winetypes" 
      :key="winetype.id" 
      :item="winetype"
      title="id"
      subtitle="description"
      :fields="[
        
        
        
        
        
        { key: 'category', label: 'Category' },
        
      ]"
      timestamp="createdAt"
      hover
    >
      <template #actions>
        <BaseCrudActions 
          structure="winetypes"
          :item="winetype" 
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
          @click="$emit('view', winetype)"
        >
          View Details
        </UButton>
      </template>
    </WineTypeGridCard>
  </div>
</template>

<script setup lang="ts">
import type { WineType } from '../stores/winetype'
import { computed } from 'vue'

const props = defineProps<{
  winetype: WineType[]
}>()

const winetypes = computed(() => props.winetype)

defineEmits<{
  (e: 'view', winetype: WineType): void
  (e: 'edit', winetype: WineType): void
  (e: 'delete', winetype: WineType): void
}>()
</script>
