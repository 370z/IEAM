
import env from '@/env';
// import firebase from 'firebase/app'
import firebase from 'firebase'
import 'firebase/storage'
// import 'firebase/firestore' // <- add

if (!firebase.apps.length) {
  firebase.initializeApp(env.FIREBASE)
}
const storage = firebase.storage()

export { storage }


