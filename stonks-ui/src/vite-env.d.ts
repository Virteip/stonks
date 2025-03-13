/// <reference types="vite/client" />

declare module '*.vue' {
    import type { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
  }
  
  declare module 'vue' {
    export type Ref<T> = import('vue').Ref<T>
    export function ref<T>(value: T): Ref<T>
    export function ref<T>(): Ref<T | undefined>
    export function computed<T>(getter: () => T): Ref<T>
    export function reactive<T extends object>(target: T): T
    export function onMounted(callback: () => void): void
    export function defineProps<T>(): T
    export function defineEmits<T extends string[]>(...events: T): { [K in T[number]]: (...args: any[]) => void }
    
    // App factory
    export interface App<HostElement = any> {
      mount(rootContainer: HostElement | string): ComponentPublicInstance
      use<Options extends any[]>(plugin: Plugin<Options>, ...options: Options): this
      component(name: string): Component | undefined
      component(name: string, component: Component): this
      directive(name: string): Directive | undefined
      directive(name: string, directive: Directive): this
      provide<T>(key: InjectionKey<T> | string, value: T): this
      config: AppConfig
      [key: string]: any
    }
    
    export type CreateAppFunction<HostElement> = (rootComponent: Component, rootProps?: Data | null) => App<HostElement>
    
    export function createApp(rootComponent: Component, rootProps?: Data | null): App
  }