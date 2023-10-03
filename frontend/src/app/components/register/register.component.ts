import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  registrationForm: FormGroup;
  formErrors: any = {};

  constructor(
    private fb: FormBuilder,
    private http: HttpClient,
    private router: Router // Inject Router
  ) {
    this.registrationForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', [Validators.required]]
    });
  }

  onSubmit() {
    const formData = this.registrationForm.value;
    if (formData.password !== formData.confirmPassword) {
      this.formErrors.confirmPassword = 'Passwords do not match';
      return;
    }
    if (this.registrationForm.valid) {
      // Clear any previous form errors
      this.formErrors = {};

      this.http.post('http://localhost:11001/api/membership/signUp', formData)
        .pipe(
          catchError(this.handleError)
        )
        .subscribe(
          (response) => {
            // Handle success response here
            console.log('Registration successful:', response);
            // Store the JWT token (adjust the property name as per your server response)
            const token = response.oauthAccessToken; // Replace 'token' with the actual property name

            // Store the token in localStorage (you can use a more secure storage method)
            localStorage.setItem('token', token);

            // Navigate to the home path
            this.router.navigate(['/home']); 
            // You can also navigate to a login page or show a success message to the user
          },
          (error) => {
            // Handle server-side validation errors here
            if (error.status === 400 && error.error) {
              this.formErrors = error.error;
            } else {
              this.formErrors = error.error;
              this.handleError(error);
            }
          }
        );
    } else {
      // Mark form fields as touched to trigger validation error messages
      this.markFormFieldsAsTouched(this.registrationForm);
    }
  }

  private handleError(error: any): Observable<any> {
    console.error('Registration error:', error);
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
