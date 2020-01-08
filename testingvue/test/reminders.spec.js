import { mount } from 'vue-test-utils'
import expect from 'expect'
import Reminders from '../src/components/Reminders.vue'

describe ('Reminders', () => {
  let wrapper
  
  beforeEach(() => {
    wrapper = mount(Reminders)
  })
  
  it ('hides the reminders list if there are none', () => {
    let wrapper = mount(Reminders)
    
    expect(wrapper.contains('ul')).toBe(false)
  })
  
  it ('can add reminders', () => {
    let wrapper = mount(Reminders)
    
    addReminder('Go to the store')
    
    expect(remindersList()).toContain('Go to the store')
    expect(remindersList()).toContain('Finish screencast')
  })
  
  it ('can remove any reminder', () => {
    addReminder('Go to the store')
    addReminder('Finish screencast')
    
    let deleteButton = wrapper.find('ul > li:fist-child .remove')
    
    deleteButton.trigger('click')
    
    expect(remindersList())not.toContain('Go to the store')
  })
  
  function addReminder(body) {
      let newReminder = wrapper.find('.new-reminder')
      
      newReminder.element.value = body
      newReminder.trigger('input')
      
      wrapper.find('button').trigger('click')
    }
    
    function remindersList() {
      return wrapper.find('ul').text()
    }
})