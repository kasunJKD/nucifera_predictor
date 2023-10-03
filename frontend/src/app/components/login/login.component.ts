import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm: FormGroup;
  formErrors: any = {};

  constructor(
    private fb: FormBuilder,
    private http: HttpClient,
    private router: Router // Inject Router
  ) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  onSubmit() {
    if (this.loginForm.valid) {
      const formData = this.loginForm.value;

      // Clear any previous form errors
      this.formErrors = {};

      this.http.post('http://localhost:11001/api/membership/passwordSignIn', formData)
        .pipe(
          catchError(this.handleError)
        )
        .subscribe(
          (response: any) => {
            // Handle success response here
            console.log('Login successful:', response);

            // Store the JWT token (adjust the property name as per your server response)
            const token = response.token; // Replace 'token' with the actual property name

            // Store the token in localStorage (you can use a more secure storage method)
            localStorage.setItem('token', token);

            // Redirect to the home path or any other desired route
            this.router.navigate(['/home']); // Replace '/home' with your desired path
          },
          (error) => {
            // Handle login error messages here
            if (error.status === 400 && error.error) {
              this.formErrors = error.error;
            } else {
              this.formErrors = error.error;
              console.log(this.formErrors)
              this.handleError(error);
            }
          }
        );
    } else {
      // Mark form fields as touched to trigger validation error messages
      this.markFormFieldsAsTouched(this.loginForm);
    }
  }

  private handleError(error: any): Observable<any> {
    console.error('Login error:', error);
    return throwError(error);
  }

  markFormFieldsAsTouched(formGroup: FormGroup) {
    Object.values(formGroup.controls).forEach(control => {
      control.markAsTouched();

      if (control instanceof FormGroup) {
        this.markFormFieldsAsTouched(control);
      }
    });
  }
}