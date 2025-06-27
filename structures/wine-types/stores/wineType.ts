import type { BaseItem } from '@@/app/stores/baseStore'

/**
 * WineType entity interface
 */
export interface WineType extends BaseItem {
  name?: string;
  description?: string;
  category?: string;
  createdAt: string;
  updatedAt: string;
}

/**
 * Utility functions for WineType entity
 */
export class WineTypeModel {
  /**
   * Factory method to create a WineType from JSON data
   * Handles conversion from snake_case to camelCase
   */
  static fromJson(json: Record<string, any>): WineType {
    // Helper function to convert camelCase to snake_case
    const toSnakeCase = (str: string) => str.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`)
    
    // Helper function to get value from either camelCase or snake_case key
    const getValue = (obj: Record<string, any>, camelKey: string) => {
      if (obj[camelKey] !== undefined) return obj[camelKey]
      const snakeKey = toSnakeCase(camelKey)
      return obj[snakeKey]
    }
    
    return {
      id: Number(json.id),
      name: getValue(json, 'name'),
      description: getValue(json, 'description'),
      category: getValue(json, 'category'),
      createdAt: getValue(json, 'createdAt'),
      updatedAt: getValue(json, 'updatedAt')
    }
  }

  /**
   * Factory method to create a list of WineType from JSON data
   */
  static fromJsonList(jsonList: Record<string, any>[]): WineType[] {
    return jsonList.map(json => this.fromJson(json))
  }

  /**
   * Convert WineType to JSON format (for API requests)
   * Optionally converts camelCase to snake_case
   */
  static toJson(entity: Partial<WineType>, useSnakeCase: boolean = false): Record<string, any> {
    if (!useSnakeCase) {
      return { ...entity }
    }
    
    // Convert to snake_case for API
    const result: Record<string, any> = {}
    
    Object.entries(entity).forEach(([key, value]) => {
      // Convert camelCase to snake_case
      const snakeKey = key.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`)
      result[snakeKey] = value
    })
    
    return result
  }
}
