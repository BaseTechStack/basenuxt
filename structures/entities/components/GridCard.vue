<!-- Generic GridCard component for entity -->
<template>
  <UCard 
    :class="[
      'grid-card',
      `grid-card--${size}`,
      { 'grid-card--hover': hover }
    ]" 
    :ui="{
      body: 'p-0',
      header: 'p-4',
      footer: 'p-4'
    }"
  >
    <template #header>
      <slot name="header">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-lg font-semibold line-clamp-1">{{ computedTitle }}</h3>
            <p v-if="computedSubtitle" class="text-sm text-gray-500 dark:text-gray-400 line-clamp-1">{{ computedSubtitle }}</p>
          </div>
          <slot name="actions"></slot>
        </div>
      </slot>
    </template>
    
    <div class="grid-card-content p-4">
      <slot name="content">
        <div class="space-y-3">
          <div 
            v-for="(field, index) in displayFields" 
            :key="index"
            class="flex items-start gap-2"
          >
            <UIcon 
              v-if="field.icon" 
              :name="field.icon" 
              class="mt-0.5 flex-shrink-0 text-gray-400 dark:text-gray-500" 
            />
            <div>
              <div class="text-xs text-gray-500 dark:text-gray-400">{{ field.label }}</div>
              <div class="line-clamp-2">{{ formatFieldValue(field.key, getFieldValue(field.key), field.format) }}</div>
            </div>
          </div>
        </div>
      </slot>
    </div>
    
    <template #footer>
      <slot name="footer">
        <div class="flex items-center justify-between">
          <div v-if="timestamp" class="text-xs text-gray-500 dark:text-gray-400">
            {{ formatDate(getFieldValue(timestamp)) }}
          </div>
          <div class="flex items-center gap-2">
            <slot name="footer-actions"></slot>
          </div>
        </div>
      </slot>
    </template>
  </UCard>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { format, formatDistance } from 'date-fns'

interface Field {
  key: string
  label: string
  icon?: string
  format?: string
}

interface Props {
  item: Record<string, any>
  title: string
  subtitle?: string
  fields?: Field[]
  timestamp?: string
  size?: 'sm' | 'md' | 'lg'
  hover?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  size: 'md',
  hover: true,
  fields: () => []
});

// Computed properties
const displayFields = computed(() => {
  return props.fields;
});

const computedTitle = computed(() => {
  // First check if the item has a name property
  if (props.item?.name) {
    return props.item.name;
  }
  // Then check if the item has a title property
  if (props.item?.title) {
    return props.item.title;
  }
  // Finally use the specified title field
  return getFieldValue(props.title);
});

const computedSubtitle = computed(() => {
  if (!props.subtitle) return '';
  return getFieldValue(props.subtitle);
});

// Methods
function getFieldValue(key: string): any {
  return props.item?.[key] ?? '';
}

function formatFieldValue(key: string, value: any, format?: string): string {
  if (value === undefined || value === null) return '—';
  
  if (format === 'date') {
    return formatDate(value);
  }
  
  if (format === 'currency') {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(value);
  }
  
  return String(value);
}

function formatDate(date: string | Date): string {
  try {
    const dateObj = typeof date === 'string' ? new Date(date) : date;
    
    // Check if date is valid
    if (isNaN(dateObj.getTime())) {
      return '—';
    }
    
    // Return relative time for recent dates (within last 7 days)
    const now = new Date();
    const diff = now.getTime() - dateObj.getTime();
    if (diff < 7 * 24 * 60 * 60 * 1000) {
      return formatDistance(dateObj, now, { addSuffix: true });
    }
    
    return format(dateObj, 'MMM d, yyyy');
  } catch {
    return '—';
  }
}
</script>

<style scoped>
.grid-card {
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.grid-card--hover:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.grid-card--sm .grid-card-content {
  max-height: 100px;
  overflow: hidden;
}

.grid-card--md .grid-card-content {
  max-height: 150px;
  overflow: hidden;
}

.grid-card--lg .grid-card-content {
  max-height: 200px;
  overflow: hidden;
}
</style>
