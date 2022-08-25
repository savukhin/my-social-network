import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppearanceService } from 'src/app/services/backend-api/appearance.service';
import { AuthService } from 'src/app/services/backend-api/auth.service';
import { ContentService } from 'src/app/services/backend-api/content.service';
import { Post } from 'src/models/post';
import { User, UserCompressed } from 'src/models/user';
import { UserPage } from 'src/models/UserPage';

@Component({
    selector: 'app-user-page',
    templateUrl: './user-page.component.html',
    styleUrls: ['./user-page.component.scss']
})
export class UserPageComponent implements AfterViewInit {
    user?: User;
    profile = new UserPage()
    editStatus = false
    avatarEditing = false
    photoUploadIsSet = false
    imageSrc = ""
    
    @ViewChild('pleaseDoIt') input: ElementRef<HTMLInputElement> = {} as ElementRef;
    @ViewChild('postText') postText: ElementRef<HTMLTextAreaElement> = {} as ElementRef;
    @ViewChild('avatarInput') avatarInput: ElementRef<HTMLInputElement> = {} as ElementRef;

    showEditStatus() {
        setTimeout(() => {
            this.input.nativeElement.focus();
            this.input.nativeElement.value = this.profile.status;
        }, 0)
        this.editStatus = true;
    }

    hideEditStatus() {
        this.editStatus = false;
    }

    constructor(private route: ActivatedRoute, private router: Router, private auth: AuthService, private cdref: ChangeDetectorRef, private content: ContentService, private appearance: AppearanceService) {
        if (this.auth.userSubscription == undefined) 
            return
        
        this.auth.userSubscription.subscribe((response) => {
            if (response == false)
                return;
            this.user = response
            this.cdref.detectChanges()
        })

        this.appearance.blackout.subscribe(val => {
            if (val == false) {
                this.avatarEditing = false
                this.photoUploadIsSet = false
                this.imageSrc = ""
            }
        })
    }

    editStatusClick(): void {
        const newStatus = this.input.nativeElement.value
        const subscription = this.auth.changeProfileStatus(newStatus)
        if (subscription != false) {
            subscription.subscribe((response) => {
                if (response != false) {
                    this.profile.status = newStatus
                    this.hideEditStatus()
                }
            })
        }
    }

    createPostClick(): void {
        const text = this.postText.nativeElement.value
        const subscription = this.content.createPosts(text)
        if (subscription == false) {
            location.reload()
            return
        }

        subscription.subscribe(() => {
            location.reload()
        })
    }

    likePostClick(post: Post): void {
        const subscription = this.content.toggleLikePost(post.id)
        if (subscription == false) {
            location.reload()
            return
        }

        subscription.subscribe(() => {
            post.has_current_user_like = !post.has_current_user_like;
            post.current_likes += (post.has_current_user_like? 1 : -1)
        })
    }

    addFriendClick(): void {
        if (this.profile.id == this.user?.id)
            return

        const subscription = this.auth.addToFriends(this.profile.id)
        if (subscription == false) {
            return
        }

        subscription.subscribe((response) => {
            if (response != false) {
                this.profile.added_to_friends = true
                location.reload()
            }
        })
    }

    deleteFriendClick(): void {
        if (this.profile.id == this.user?.id)
            return

        const subscription = this.auth.deleteFriend(this.profile.id)
        if (subscription == false) {
            return
        }

        subscription.subscribe((response) => {
            if (response != false) {
                this.profile.added_to_friends = true
                location.reload()
            }
        })
    }

    navigateToFriend(friend: UserCompressed) {
        this.router.navigate([`/user`, friend.id]).then(() => {
            location.reload()
        })
    }

    changeAvatarClick() {
        this.avatarEditing = true
        this.appearance.blackout.next(true)
    }

    onChangeUpload(event: Event) {
        let target = (event.target as HTMLInputElement)
        if (target.files && target.files[0]) {
            const file = target.files[0];
    
            const reader = new FileReader();
            reader.onload = e => this.imageSrc = reader.result as string;
    
            reader.readAsDataURL(file);
            this.photoUploadIsSet = true
        }
    }

    uploadClick() {
        let target = this.avatarInput.nativeElement
        if (!target.files || !target.files[0]) {
            return
        }

        const file = target.files[0];

        let subscription = this.auth.changeProfileAvatar(file)
        if (subscription == false)
            return
        
        subscription.subscribe((response) => {
            console.log(response);
            // if (response != false)
            //     location.reload()
        })
    }

    ngAfterViewInit(): void {
        let id = this.route.snapshot.paramMap.get("id");
        if (id == null) {
            this.auth.userSubscription?.subscribe((response) => {
                if (response != false) {
                    this.user = response
                    this.router.navigateByUrl('user/' + this.user.id)
                }
            })

            return
        }

        const subscription = this.auth.getProfile(+id)
        if (subscription != false) {
            subscription.subscribe(
                response => {
                    if (response == false)
                        return 

                    this.profile = response
                    console.log(this.profile);
                    
                    this.cdref.detectChanges();

                    this.content.getUserPosts(this.profile.id).subscribe((response) => {
                        if (response == false)
                            return;
                        
                        this.profile.posts = response
                        this.cdref.detectChanges()
                    })

                    // this.auth.GetProfileAvatar(this.profile.id).subscribe((response) => {

                    // })
                }
            )
        }
        
        // })
    }
}
