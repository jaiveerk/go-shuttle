package listener

import (
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-shuttle/common/options/listeneropts"
	"time"
)

type Option = listeneropts.Option

type ManagementOption = listeneropts.ManagementOption

func WithMessageLockAutoRenewal(interval time.Duration) Option {
	return listeneropts.WithMessageLockAutoRenewal(interval)
}

// WithConnectionString configures a listener with the information provided in a Service Bus connection string
func WithConnectionString(connStr string) ManagementOption {
	return listeneropts.WithConnectionString(connStr)
}

// WithEnvironmentName configures the azure environment used to connect to Servicebus. The environment value used is
// then provided by Azure/go-autorest.
// ref: https://github.com/Azure/go-autorest/blob/c7f947c0610de1bc279f76e6d453353f95cd1bfa/autorest/azure/environments.go#L34
func WithEnvironmentName(environmentName string) ManagementOption {
	return listeneropts.WithEnvironmentName(environmentName)
}

// WithManagedIdentityClientID configures a listener with the attached managed identity and the Service bus resource name
func WithManagedIdentityClientID(serviceBusNamespaceName, managedIdentityClientID string) ManagementOption {
	return listeneropts.WithManagedIdentityClientID(serviceBusNamespaceName, managedIdentityClientID)
}

// WithToken configures a listener with a AAD token
func WithToken(serviceBusNamespaceName string, spt *adal.ServicePrincipalToken) ManagementOption {
	return listeneropts.WithToken(serviceBusNamespaceName, spt)
}

// WithManagedIdentityResourceID configures a listener with the attached managed identity and the Service bus resource name
func WithManagedIdentityResourceID(serviceBusNamespaceName, managedIdentityResourceID string) ManagementOption {
	return listeneropts.WithManagedIdentityResourceID(serviceBusNamespaceName, managedIdentityResourceID)
}

// WithDetails allows listeners to control Queue details for longer lived operations.
// If you using RetryLater you probably want this. Passing zeros leaves it up to Service bus defaults
func WithSubscriptionDetails(lock time.Duration, maxDelivery int32) ManagementOption {
	return listeneropts.WithDetails(lock, maxDelivery)
}

// WithLockDuration allows listeners to control LockDuration. Passing zeros leaves it up to Service bus defaults
func WithLockDuration(lock time.Duration) ManagementOption {
	return listeneropts.WithLockDuration(lock)
}

// WithQueueMaxDeliveryCount allows listeners to control MaxDeliveryCount. Passing zeros leaves it up to Service bus defaults
func WithMaxDeliveryCount(maxDelivery int32) ManagementOption {
	return listeneropts.WithMaxDeliveryCount(maxDelivery)
}

// WithPrefetchCount the receiver to quietly acquires more messages, up to the PrefetchCount limit. A single Receive call to the ServiceBus api
// therefore acquires several messages for immediate consumption that is returned as soon as available.
// Please be aware of the consequences : https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-prefetch#if-it-is-faster-why-is-prefetch-not-the-default-option
func WithPrefetchCount(prefetch uint32) Option {
	return listeneropts.WithPrefetchCount(prefetch)
}

func WithMaxConcurrency(concurrency int) Option {
	return listeneropts.WithMaxConcurrency(concurrency)
}
