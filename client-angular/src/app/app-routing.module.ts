import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { ChatComponent } from './chat/chat.component';
import { FriendsComponent } from './friends/friends.component';
import { MessagesComponent } from './messages/messages.component';
import { UpdateProfileComponent } from './profile/update-profile/update-profile.component';
import { UserPageComponent } from './profile/user-page/user-page.component';

const routes: Routes = [
    {
        path: 'user',
        component: UserPageComponent
    },
    {
      path: 'messages',
      component: MessagesComponent
    },
    {
      path: 'chat/:id',
      component: ChatComponent
    },
    {
      path: 'friends',
      component: FriendsComponent
    },
    {
      path: 'edit-profile',
      component: UpdateProfileComponent
    },
    {
      path: 'login',
      component: LoginComponent
    },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }