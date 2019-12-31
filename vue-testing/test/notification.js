import Vue from 'vue'
import test from 'ava'
import Notifications from '../src/Notifications'

test('that it renders a notification', t => {
  t.is(Notification.data().message, 'Hello World')
})