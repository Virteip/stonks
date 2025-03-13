import { defineStore } from 'pinia';
import { stockService, Stock, PaginatedStocks } from '@/services/api';

export const useStocksStore = defineStore('stocks', {
  state: () => ({
    stocks: [] as Stock[],
    loading: false,
    error: null as string | null,
    pagination: {
      totalCount: 0,
      pageSize: 10,
      currentPage: 1,
      totalPages: 0
    },
    searchTicker: ''
  }),

  actions: {
    async fetchStocks(page = 1) {
      this.loading = true;
      this.error = null;
      try {
        const data: PaginatedStocks = await stockService.getStocks(page);
        this.stocks = data.stocks;
        this.pagination = {
          totalCount: data.total_count,
          pageSize: data.page_size,
          currentPage: data.page,
          totalPages: data.total_pages
        };
      } catch (error) {
        this.error = error instanceof Error ? error.message : 'Failed to fetch stocks';
        this.stocks = [];
      } finally {
        this.loading = false;
      }
    },

    async searchStock(ticker: string) {
      this.loading = true;
      this.error = null;
      this.searchTicker = ticker;
      try {
        const stocks = await stockService.getStockByTicker(ticker);
        this.stocks = stocks;
        
        if (this.stocks.length === 0) {
          this.error = `No stocks found for ticker: ${ticker}`;
        }
      } catch (error) {
        this.error = error instanceof Error ? error.message : 'Failed to find stock';
        this.stocks = [];
      } finally {
        this.loading = false;
      }
    },

    clearSearch() {
      this.searchTicker = '';
      this.fetchStocks();
    }
  }
});