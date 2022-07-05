import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MessagesComponent } from './messages/messages.component';
import { UserPageComponent } from './user-page/user-page.component';

const routes: Routes = [
    {
        path: 'user',
        component: UserPageComponent
    },
    {
      path: 'messages',
      component: MessagesComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }