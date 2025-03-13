import axios from 'axios';

// Get environment variables with fallbacks
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '';
const API_KEY = import.meta.env.VITE_API_KEY || '';

// Create an axios instance with default configuration
const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'X-API-Key': API_KEY
  }
});

export interface Stock {
  id: string;
  ticker: string;
  company: string;
  brokerage: string;
  action: string;
  rating_from: string;
  rating_to: string;
  target_from: number;
  target_to: number;
  time: string;
}

export interface StockRecommendation {
  stock: Stock;
  score: number;
  reason: string;
}

export interface PaginatedStocks {
  stocks: Stock[];
  total_count: number;
  page_size: number;
  page: number;
  total_pages: number;
}

export const stockService = {
  async getStocks(page = 1, pageSize = 10): Promise<PaginatedStocks> {
    const response = await apiClient.get<PaginatedStocks>('/stocks', {
      params: { page, page_size: pageSize }
    });
    return response.data;
  },

  async getStockByTicker(ticker: string): Promise<Stock[]> {
    const response = await apiClient.get<Stock[]>(`/stock/${ticker}`);
    return response.data;
  }
};

export const recommendationService = {
  async getRecommendations(): Promise<StockRecommendation[]> {
    const response = await apiClient.get<StockRecommendation[]>('/recommendations');
    return response.data;
  }
};