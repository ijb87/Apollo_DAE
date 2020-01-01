import { mount } from 'vue-test-utitls'
import Counter from '../src/components/Counter.js'
import expect from 'expect'

describe ('Counter', () => {
  it ('defaults to a count of 0', () => {
    let wrapper = mount(Counter)
    
    expect(wrapper.vm.count).toBe(0)
  })
})