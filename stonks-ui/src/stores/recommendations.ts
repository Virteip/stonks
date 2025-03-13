import { defineStore } from 'pinia';
import { recommendationService, type StockRecommendation } from '@/services/api';

// Define interface for the store state
interface RecommendationsState {
  recommendations: StockRecommendation[];
  loading: boolean;
  error: string | null;
}

export const useRecommendationsStore = defineStore('recommendations', {
  state: (): RecommendationsState => ({
    recommendations: [] as StockRecommendation[],
    loading: false,
    error: null
  }),

  actions: {
    async fetchRecommendations() {
      this.loading = true;
      this.error = null;
      try {
        this.recommendations = await recommendationService.getRecommendations();
      } catch (error) {
        this.error = error instanceof Error ? error.message : 'Failed to fetch recommendations';
        console.error('Error fetching recommendations:', error);
      } finally {
        this.loading = false;
      }
    }
  }
});