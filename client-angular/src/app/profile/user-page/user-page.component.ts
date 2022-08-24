import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from 'src/app/services/backend-api/auth.service';
import { ContentService } from 'src/app/services/backend-api/content.service';
import { User } from 'src/models/user';
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
    
    @ViewChild('pleaseDoIt') input: ElementRef<HTMLInputElement> = {} as ElementRef;

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

    constructor(private route: ActivatedRoute, private router: Router, private auth: AuthService, private cdref: ChangeDetectorRef, private content: ContentService) {
        if (this.auth.userSubscription == undefined) 
            return
        
        this.auth.userSubscription.subscribe((response) => {
            if (response == false)
                return;
            this.user = response
            this.cdref.detectChanges()
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
                    if (response != false) {
                        this.profile = UserPage.FromUser(response);

                        this.content.getUserPosts(this.profile.id).subscribe((response) => {
                            if (response == false)
                                return;
                            
                            this.profile.posts = response
            
                            this.cdref.detectChanges()
                        })
                    }
                    this.cdref.detectChanges();
                }
            )
        }
        
        // })
    }
}
